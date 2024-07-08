package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gofor-little/env"
	"github.com/golang-jwt/jwt"
	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/middleware"
	"github.com/semperprimum/devlinks/server/models"
	"github.com/semperprimum/devlinks/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type GetUsersResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type GetUserResponse struct {
	ID           int            `json:"id"`
	Email        string         `json:"email"`
	DisplayEmail sql.NullString `json:"display_email"`
	FirstName    sql.NullString `json:"first_name"`
	LastName     sql.NullString `json:"last_name"`
	PicPath      sql.NullString `json:"pic_path"`
}

type LoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type CreateUserData struct {
	ID       int
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UpdateUserInfoData struct {
	DisplayEmail string `json:"display_email,omitempty" validate:"omitempty,email"`
	FirstName    string `json:"first_name,omitempty" validate:"omitempty"`
	LastName     string `json:"last_name,omitempty" validate:"omitempty"`
	Link         string `json:"link,omitempty" validate:"omitempty"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, email FROM users")

	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	users := []GetUsersResponse{}
	for rows.Next() {
		var user GetUsersResponse
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUser(r)

	var user GetUserResponse
	err := db.DB.QueryRow("SELECT id, first_name, last_name, email, pic_path, display_email FROM users WHERE id = $1", userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PicPath, &user.DisplayEmail)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WriteJSONError(w, "Bad request", http.StatusBadRequest)
		return
	}

	validationErrors := utils.ValidateStruct(user)
	if validationErrors != nil {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	err = db.DB.QueryRow("INSERT INTO users(email, password) VALUES($1, $2) RETURNING id", &user.Email, &user.Password).Scan(&user.ID)
	if user.ID == 0 {
		utils.WriteJSONError(w, "User already exists", http.StatusBadRequest)
		return
	}

	newUser := CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var data LoginData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.WriteJSONError(w, "Bad request", http.StatusBadRequest)
		return
	}

	validationErrors := utils.ValidateStruct(data)
	if validationErrors != nil {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	var user models.User
	err = db.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", data.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSONError(w, "User not found", http.StatusNotFound)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		utils.WriteJSONError(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	var token string
	token, err = GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res = LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	var data UpdateUserInfoData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.WriteJSONError(w, "Bad request", http.StatusBadRequest)
		return
	}

	validationErrors := utils.ValidateStruct(data)
	if validationErrors != nil {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	userID := middleware.GetUser(r)

	if data.FirstName != "" {
		_, err = db.DB.Exec("UPDATE users SET first_name = $1 WHERE id = $2", data.FirstName, userID)
		if err != nil {
			utils.WriteJSONError(w, "Failed updating first_name field", http.StatusInternalServerError)
			return
		}
	}

	if data.LastName != "" {
		_, err = db.DB.Exec("UPDATE users SET last_name = $1 WHERE id = $2", data.LastName, userID)
		if err != nil {
			utils.WriteJSONError(w, "Failed updating last_name field", http.StatusInternalServerError)
			return
		}
	}

	if data.DisplayEmail != "" {
		_, err = db.DB.Exec("UPDATE users SET display_email = $1 WHERE id = $2", data.DisplayEmail, userID)
		if err != nil {
			utils.WriteJSONError(w, "Failed updating display_email field", http.StatusInternalServerError)
			return
		}
	}

	if data.Link != "" {
		_, err = db.DB.Exec("UPDATE users SET link = $1 WHERE id = $2", data.Link, userID)
		if err != nil {
			utils.WriteJSONError(w, "Failed updating link field", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GenerateJWT(userID string) (string, error) {
	secret, err := env.MustGet("JWT_SECRET")
	if err != nil {
		panic(err)
	}

	// Set token expiration time
	expirationTime := time.Now().Add(120 * time.Hour)

	// Create the claims
	claims := &jwt.MapClaims{
		"id":  userID,
		"exp": expirationTime.Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
