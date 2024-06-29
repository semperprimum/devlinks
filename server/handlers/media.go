package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/middleware"
	"github.com/semperprimum/devlinks/server/utils"
)

func UploadPicture(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("File size too big: %s", err.Error()), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("Unable to get file: %s", err.Error()), http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileType := handler.Header.Get("Content-Type")
	if fileType != "image/jpeg" && fileType != "image/png" {
		utils.WriteJSONError(w, "Invalid file type. Only JPG and PNG are allowed.", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUser(r)

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("Unable to create directory: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("user_%s", userID)
	filePath := filepath.Join("uploads", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("Unable to crate file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("Unable to save file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	relativePath := "/uploads/" + fileName
	_, err = db.DB.Exec("UPDATE users SET pic_path = $1 WHERE id = $2", relativePath, userID)
	if err != nil {
		utils.WriteJSONError(w, fmt.Sprintf("Unable to update database: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile picture uploaded successfully!", "file_path": relativePath})
}
