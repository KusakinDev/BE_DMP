package loadimage

import (
	cdc "back/config/cloudinaryConfig"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	// Парсим форму с файлом
	err := r.ParseMultipartForm(10 << 20) // Ограничение: 10MB
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	// Извлекаем файл из запроса
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename) // Получаем расширение файла
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Загрузка файла в Cloudinary
	uploadResult, err := cdc.CLD.Upload.Upload(r.Context(), file, uploader.UploadParams{
		Folder:   "product_image", // Опционально: папка в Cloudinary
		PublicID: newFileName,
	})
	if err != nil {
		http.Error(w, "Ошибка загрузки в Cloudinary", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"url": "%s", "public_id": "%s"}`, uploadResult.SecureURL, uploadResult.PublicID)
}
