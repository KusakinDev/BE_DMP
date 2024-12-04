package main

import (
	cdc "back/config/cloudinaryConfig"
	dbA "back/db"
	lg "back/enters/login"
	gaf "back/feed/getAllFeed"
	uli "back/image/loadImage"
	cpc "back/sellers/createProductCard"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors" // Импортируем библиотеку CORS
)

func main() {

	// Инициализация базы данных
	dbA.InitDB()
	defer dbA.DB.Close()

	// Настройка Cloudinary
	cdc.CloudinaryConfig()

	// Создаем новый объект CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // Разрешаем запросы с фронтенда
		AllowCredentials: true,                                                // Разрешаем отправку cookies
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем различные методы
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           // Разрешаем заголовки
	})

	// Оборачиваем обработчики в CORS middleware
	mux := http.NewServeMux()

	// Настройка маршрутов
	mux.HandleFunc("/login", lg.Login)
	mux.HandleFunc("/getAllFeed", gaf.GetAllFeed)
	mux.HandleFunc("/createProductCard", cpc.CreateProductCard)
	mux.HandleFunc("/uploadProductImage", uli.UploadImage)

	// Применяем CORS middleware
	handler := c.Handler(mux)

	// Запуск сервера
	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
