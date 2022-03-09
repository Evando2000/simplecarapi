package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var carList []Car

func prepare() {
	err := os.Mkdir("./files", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}

	newCar := Car{
		ID:    "0",
		Model: "SUV",
		Color: "Black",
		Brand: "newBrand",
	}

	carList = append(carList, newCar)
}

func main() {
	prepare()
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost", "herokuapp.com"})

	carGroupRouter := router.Group("/car")
	{
		carGroupRouter.GET("/", getAllCarsHandler)
		carGroupRouter.POST("/", createCarHandler)
		carGroupRouter.GET("/:id", getCarHandler)
		carGroupRouter.PUT("/:id", updateCarHandler)
		carGroupRouter.DELETE("/:id", deleteCarHandler)

	}

	fileGroupRouter := router.Group("/file")
	{
		fileGroupRouter.POST("/", uploadFileHandler)
	}

	port := getEnv("PORT", "8080")
	router.Run(":" + port)
}