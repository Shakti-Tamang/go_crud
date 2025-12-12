package controllers

import (
	"github.com/Shakti-Tamang/crud-app/initializers"
	"github.com/Shakti-Tamang/crud-app/models"
	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {

	type body struct {
		ProductName string
		Description string
		Quantity    int16
	}

	var bodyData body

	if err := c.BindJSON(&bodyData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	var existing models.Item

	test := initializers.DB.Where("product_name = ?", bodyData.ProductName).First(&existing).Error
	if test == nil {

		c.JSON(400, gin.H{"error": "Product name already exists"})
		return
	}

	item := models.Item{ProductName: bodyData.ProductName, Description: bodyData.Description, Quantity: bodyData.Quantity}

	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.JSON(500, gin.H{"er": result.Error.Error()})
		return //stops fucntion in this point
	}

	c.JSON(200, gin.H{"message": "Item created successfully", "item": item})

}
