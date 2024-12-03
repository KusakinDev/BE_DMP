package main

import (
	dbA "back/db"
	lg "back/enters/login"
	gaf "back/feed/getAllFeed"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors" // Импортируем библиотеку CORS
)

func main() {

	dbA.InitDB()
	defer dbA.DB.Close()

	// Создаем новый объект CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Разрешаем запросы с указанного источника
		AllowCredentials: true,
	})

	// Оборачиваем обработчики в CORS middleware
	mux := http.NewServeMux()

	mux.HandleFunc("/login", lg.Login)
	mux.HandleFunc("/getAllFeed", gaf.GetAllFeed)

	handler := c.Handler(mux)

	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
