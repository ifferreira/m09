package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context, file *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(file.Filename)
	
	filename := uuid.New().String() + ext
	
	dirPath := filepath.Join("uploads", folder)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	
	filePath := filepath.Join(dirPath, filename)
	
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}
	
	return "/" + filePath, nil
}

func ValidateFileType(file *multipart.FileHeader, allowedTypes []string) error {
	contentType := file.Header.Get("Content-Type")
	
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			return nil
		}
	}
	
	return fmt.Errorf("invalid file type: %s", contentType)
}

func DeleteFile(filePath string) error {
	cleanPath := strings.TrimPrefix(filePath, "/")
	
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}
	
	return os.Remove(cleanPath)
} 
