package main

import (
	authmiddleware "back/auth/authMiddleware"
	corsmiddleware "back/auth/corsMiddleware"
	refreshjwt "back/auth/refreshJWT"
	cdc "back/config/cloudinaryConfig"
	dbA "back/db"
	"back/enters/login"
	getallfeed "back/feed/getAllFeed"
	loadimage "back/image/loadImage"
	createproductcard "back/sellers/createProductCard"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// Инициализация базы данных
	dbA.InitDB()
	defer dbA.DB.Close()

	// Настройка Cloudinary
	cdc.CloudinaryConfig()

	r := gin.Default()

	r.Use(corsmiddleware.CorsMiddleware())

	// Настройка маршрутов
	r.POST("/login", login.Login)
	r.POST("/refresh", refreshjwt.RefreshToken)

	protected := r.Group("/protected")
	protected.Use(authmiddleware.AuthMiddleware())

	protected.GET("/getAllFeed", getallfeed.GetAllFeed)
	protected.POST("/createProductCard", createproductcard.CreateProductCard)
	protected.POST("/uploadProductImage", loadimage.UploadImage)

	//mux.HandleFunc("/getCart", getcart.GetCart)

	// Запуск сервера
	log.Println("Server starting at :8080")
	log.Fatal(r.Run(":8080"))
}
