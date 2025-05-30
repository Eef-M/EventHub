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
	if oldAvatarURL != "" {
		parts := strings.Split(oldAvatarURL, "/uploads/avatars/")
		if len(parts) == 2 {
			oldFilename := parts[1]
			if oldFilename != "default_avatar.png" {
				oldPath := filepath.Join("uploads", "avatars", oldFilename)
				_ = os.Remove(oldPath)
			}
		}
	}

	uploadDir := filepath.Join("uploads", "avatars")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	newFilename := fmt.Sprintf("%s_%s", uuid.New().String(), file.Filename)
	uploadPath := filepath.Join(uploadDir, newFilename)

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		return "", err
	}

	baseURL := initializers.BaseURL
	return fmt.Sprintf("%s/uploads/avatars/%s", baseURL, newFilename), nil
}
