package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func remove(slice []Car, s int) []Car {
	return append(slice[:s], slice[s+1:]...)
}

func getCarByID(id string) (*Car, error) {
	for _, car := range carList {
		if car.ID == id {
			return &car, nil
		}
	}
	return nil, errors.New("car not found")
}

func reqValidator(c *gin.Context) (*CarRequest, error) {
	model := c.PostForm("model")
	if model == "" {
		return nil, errors.New("model invalid")
	}

	color := c.PostForm("color")
	if color == "" {
		return nil, errors.New("color invalid")
	}

	brand := c.PostForm("brand")
	if brand == "" {
		return nil, errors.New("brand invalid")
	}

	newReq := CarRequest{
		Model: model,
		Color: color,
		Brand: brand,
	}

	return &newReq, nil
}
