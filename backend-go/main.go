package main

import (
	"IllalRajinCoding/backend-api/config"
	"IllalRajinCoding/backend-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv() //memanggil fungsi LoadEnv untuk memuat file .env
	//inisialisasi database
	database.InitDB()
	//inisialiasai Gin
	router := gin.Default()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "ror!",
		})
	})

	//mulai server dengan port 3000
	router.Run(":"+ config.GetEnv("PORT", "3000"))
}