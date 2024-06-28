package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/middleware"
	"github.com/semperprimum/devlinks/server/models"
	"github.com/semperprimum/devlinks/server/utils"
)

type AddLinkRequest struct {
	UserID   int    `json:"user_id"`
	URL      string `json:"url" validate:"required"`
	Platform string `json:"platform" validate:"required"`
	TempID   string `json:"temp_id"`
}

type UpdateOrderRequest struct {
	UserID  int   `json:"user_id"`
	LinkIDs []int `json:"link_ids" validate:"dive,number"`
}

type DeleteLinkRequest struct {
	UserID int `json:"user_id"`
	LinkID int `json:"link_id"`
}

type AddLinkResponse struct {
	ID     int    `json:"id"`
	TempID string `json:"temp_id,omitempty"`
}

type UserInfo struct {
	ID           int            `json:"id"`
	DisplayEmail sql.NullString `json:"display_email"`
	FirstName    sql.NullString `json:"first_name"`
	LastName     sql.NullString `json:"last_name"`
	PicPath      sql.NullString `json:"pic_path"`
}

type GetUserLinksResponse struct {
	Links []models.Link `json:"links"`
	User  UserInfo      `json:"user"`
}

func GetUserLinks(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	rows, err := db.DB.Query("SELECT id, url, platform, order_index FROM links WHERE user_id = $1 ORDER BY order_index", userId)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	links := []models.Link{}

	for rows.Next() {
		var link models.Link
		if err := rows.Scan(&link.ID, &link.URL, &link.Platform, &link.OderIndex); err != nil {
			utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		links = append(links, link)
	}

	var user UserInfo
	err = db.DB.QueryRow("SELECT id, first_name, last_name, display_email, pic_path FROM users WHERE id = $1", userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.DisplayEmail, &user.PicPath)

	var res = GetUserLinksResponse{
		Links: links,
		User:  user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func AddLink(w http.ResponseWriter, r *http.Request) {
	var req AddLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validationErrors := utils.ValidateStruct(req); validationErrors != nil {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	if userId, err := strconv.Atoi(middleware.GetUser(r)); err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		req.UserID = userId
	}

	var maxOrderInt int
	err := db.DB.QueryRow("SELECT COALESCE(MAX(order_index), 0) FROM links WHERE user_id = $1", req.UserID).Scan(&maxOrderInt)
	if err != nil && err != sql.ErrNoRows {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var linkID int

	err = db.DB.QueryRow("INSERT INTO links (user_id, url, platform, order_index) VALUES ($1, $2, $3, $4) RETURNING id",
		req.UserID, req.URL, req.Platform, maxOrderInt+1).Scan(&linkID)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := AddLinkResponse{
		ID:     linkID,
		TempID: req.TempID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func UpdateLinksOrder(w http.ResponseWriter, r *http.Request) {
	var req UpdateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validationErrors := utils.ValidateStruct(req); validationErrors != nil {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	if userId, err := strconv.Atoi(middleware.GetUser(r)); err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		req.UserID = userId
	}

	tx, err := db.DB.Begin()
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for i, linkID := range req.LinkIDs {
		_, err = tx.Exec("UPDATE links SET order_index = $1 WHERE id = $2 AND user_id = $3", i+1, linkID, req.UserID)
		if err != nil {
			utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	var userID int
	linkID := r.PathValue("id")

	if id, err := strconv.Atoi(middleware.GetUser(r)); err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		userID = id
	}

	var orderIndex int
	err := db.DB.QueryRow("SELECT order_index FROM links WHERE id = $1 AND user_id = $2", linkID, userID).Scan(&orderIndex)
	if err == sql.ErrNoRows {
		utils.WriteJSONError(w, "Link not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM links WHERE id = $1 AND user_id = $2", linkID, userID)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE links SET order_index = order_index - 1 WHERE user_id = $1 AND order_index > $2", userID, orderIndex)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
