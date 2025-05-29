package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	ext := filepath.Ext(file.Filename)
	filename := "image-" + time.Now().Format("20060102150405") + ext
	dst := "uploads/" + filename

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}

	fullURL := initializers.BaseURL + "/" + dst
	return fullURL, nil
}

func SaveAvatarImage(c *gin.Context, file *multipart.FileHeader, oldAvatarURL string) (string, error) {
	if oldAvatarURL != "" && strings.Contains(oldAvatarURL, "/uploads/avatars/") {
		parts := strings.Split(oldAvatarURL, "/uploads/avatars/")
		if len(parts) == 2 {
			filename := parts[1]
			if filename != "default_avatar.png" {
				localPath := filepath.Join("uploads", "avatars", filename)
				_ = os.Remove(localPath)
			}
		}
	}

	uploadDir := filepath.Join("uploads", "avatars")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s_%s", uuid.New().String(), file.Filename)
	uploadPath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		return "", err
	}

	baseURL := os.Getenv("APP_BASE_URL")
	return fmt.Sprintf("%s/%s", baseURL, uploadPath), nil
}
