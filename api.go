package main

import (
	authmiddleware "back/auth/authMiddleware"
	corsmiddleware "back/auth/corsMiddleware"
	refreshjwt "back/auth/refreshJWT"
	addtocart "back/cart/addToCart"
	getcart "back/cart/getCart"
	removefromcart "back/cart/removeFromCart"
	cdc "back/config/cloudinaryConfig"
	dbA "back/db"
	"back/enters/login"
	getallfeed "back/feed/getAllFeed"
	loadimage "back/image/loadImage"
	getmygoods "back/profile/getMyGoods"
	getprofile "back/profile/getProfile"
	createproductcard "back/sellers/createProductCard"
	disableproductcard "back/sellers/disableProductCard"
	enableproductcard "back/sellers/enableProductCard"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// Инициализация базы данных
	dbA.InitDB()
	dbA.Migration()

	// Настройка Cloudinary
	cdc.CloudinaryConfig()

	r := gin.Default()

	r.Use(corsmiddleware.CorsMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Настройка маршрутов
	r.POST("/login", login.Login)
	r.POST("/refresh", refreshjwt.RefreshToken)

	protected := r.Group("/protected")

	protected.Use(authmiddleware.AuthMiddleware())

	protected.GET("/getAllFeed", getallfeed.GetAllFeed)
	protected.GET("/getCart", getcart.GetCart)
	protected.GET("/getProfile", getprofile.GetProfile)
	protected.GET("/getMyGoods", getmygoods.GetMyGoods)

	protected.POST("/uploadProductImage", loadimage.UploadImage)
	protected.POST("/createProductCard", createproductcard.CreateProductCard)
	protected.POST("/enableProductCard", enableproductcard.EnableProductCard)
	protected.POST("/disableProductCard", disableproductcard.DisableProductCard)
	protected.POST("/addToCart", addtocart.AddToCart)
	protected.POST("/removeFromCart", removefromcart.RemoveFromCart)

	// Запуск сервера
	log.Println("Server starting at :8080")
	log.Fatal(r.Run(":8080"))
}
