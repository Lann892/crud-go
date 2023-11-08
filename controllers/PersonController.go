package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func PersonCreate(c *gin.Context) {
	body := PersonRequestBody{}

	c.BindJSON(&body)

	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := config.DB.Create(&person)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(200, &person)
}

func PersonGet(c *gin.Context) {
	var persons []models.Person
	config.DB.Find(&persons)
	c.JSON(http.StatusOK, &persons)
}

func PersonGetById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	config.DB.First(&person, id)
	c.JSON(http.StatusOK, &person)
}

func PersonUpdate(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	config.DB.First(&person, id)

	body := PersonRequestBody{}

	c.BindJSON(&body)

	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := config.DB.Model(&person).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to update data"})
		return
	}

	c.JSON(200, &person)
}

func PersonDelete(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	config.DB.Delete(&person, id)
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
