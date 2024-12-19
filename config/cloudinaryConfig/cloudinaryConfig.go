package cloudinaryconfig

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

// Объявляем глобальную переменную для Cloudinary клиента
var CLD *cloudinary.Cloudinary

func CloudinaryConfig() {
	// Настраиваем Cloudinary
	var err error
	CLD, err = cloudinary.NewFromParams("djm1rjwcp", "693765952682449", "hFoMu5uLbs529mU1KY5n2hIx7vA")
	if err != nil {
		log.Fatalf("Ошибка настройки Cloudinary: %v", err) // Используем log.Fatalf для логирования ошибки
		return
	}
}
