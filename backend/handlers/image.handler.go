package handlers

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
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

	return filename, nil
}
