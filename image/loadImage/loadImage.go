package loadimage

import (
	cdc "back/config/cloudinaryConfig"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	// Парсим форму с файлом
	err := c.Request.ParseMultipartForm(10 << 20) // Ограничение: 10MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка парсинга формы"})
		return
	}

	// Извлекаем файл из запроса
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения файла"})
		return
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename) // Получаем расширение файла
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Загрузка файла в Cloudinary
	uploadResult, err := cdc.CLD.Upload.Upload(c.Request.Context(), file, uploader.UploadParams{
		Folder:   "product_image", // Опционально: папка в Cloudinary
		PublicID: newFileName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки в Cloudinary"})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"url":       uploadResult.SecureURL,
		"public_id": uploadResult.PublicID,
	})
}
