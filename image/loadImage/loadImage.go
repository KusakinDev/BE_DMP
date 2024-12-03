package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Функция для загрузки изображения в Cloudinary
func uploadImageToCloudinary(filePath string) (string, error) {
	// Создание конфигурации для Cloudinary
	cld, err := cloudinary.NewFromParams("your_cloud_name", "your_api_key", "your_api_secret")
	if err != nil {
		return "", err
	}

	// Загрузка изображения на Cloudinary
	uploadResult, err := cld.Upload.UploadFile(filePath, uploader.UploadParams{Folder: "images"})
	if err != nil {
		return "", err
	}

	// Возврат URL изображения
	return uploadResult.URL, nil
}

// Обработчик для загрузки изображения
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Получение изображения из формы
	r.ParseMultipartForm(10 << 20) // Максимум 10 MB
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Временное сохранение файла
	tempFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		http.Error(w, "Ошибка при сохранении файла", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Копирование содержимого в временный файл
	_, err = tempFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Ошибка при копировании файла", http.StatusInternalServerError)
		return
	}

	// Загрузка изображения на Cloudinary
	imageURL, err := uploadImageToCloudinary(tempFile.Name())
	if err != nil {
		http.Error(w, "Ошибка при загрузке изображения в Cloudinary", http.StatusInternalServerError)
		return
	}

	// Отправка URL изображения обратно
	fmt.Fprintf(w, "Изображение загружено! URL: %s", imageURL)
}
