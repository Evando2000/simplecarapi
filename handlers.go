package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllCarsHandler(c *gin.Context) {
	c.JSON(200, gin.H{"response": carList})
}

func deleteAllCarHandler(c *gin.Context) {
	carList = nil
	carList = []Car{}
	newCar := Car{
		ID:    "0",
		Model: "SUV",
		Color: "Black",
		Brand: "newBrand",
	}

	carList = append(carList, newCar)
	c.JSON(200, gin.H{"response": carList})
}

func getCarHandler(c *gin.Context) {
	id := c.Param("id")
	car, err := getCarByID(id)
	if err != nil {
		c.JSON(404, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": car})
}

func createCarHandler(c *gin.Context) {
	newReq, err := reqValidator(c)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	lastCar := carList[len(carList)-1]
	lastId := 0
	if i, err := strconv.Atoi(lastCar.ID); err == nil {
		lastId = i
	}
	newID := int(lastId) + 1

	newCar := Car{
		ID:    fmt.Sprintf("%d", newID),
		Model: newReq.Model,
		Color: newReq.Color,
		Brand: newReq.Brand,
	}

	carList = append(carList, newCar)
	c.JSON(201, gin.H{"response": newCar})
}

func updateCarHandler(c *gin.Context) {
	newReq, err := reqValidator(c)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	id := c.Param("id")
	car, err := getCarByID(id)
	if err != nil {
		c.JSON(404, gin.H{"response": err.Error()})
		return
	}

	car.Model = newReq.Model
	car.Color = newReq.Color
	car.Brand = newReq.Brand

	for idx, car := range carList {
		if car.ID == id {
			carList[idx].Model = newReq.Model
			carList[idx].Color = newReq.Color
			carList[idx].Brand = newReq.Brand
		}
	}
	c.JSON(200, gin.H{"response": car})
}

func deleteCarHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := getCarByID(id)
	if err != nil {
		c.JSON(404, gin.H{"response": err.Error()})
		return
	}

	for idx, car := range carList {
		if car.ID == id {
			carList = remove(carList, idx)
		}
	}
	c.JSON(200, gin.H{"response": carList})
}

func uploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	path := "files/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": path})
}
