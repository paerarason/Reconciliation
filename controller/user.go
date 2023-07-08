package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paerarason/Reconciliation/db"
)

func OrderCreate(c *gin.Context) {
	DB := db.ConnectDB()
	var order db.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.DeletedAt = nil
	existingOrder := db.Order{}
	if err := DB.Where("phoneNumber=? AND email=?", order.PhoneNumber, order.Email).First(&existingOrder).Error; err != nil {
		if err := DB.Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create order"})
			return
		}
		c.JSON(200, order)
	}
	order.LinkedId = &existingOrder.Id
	order.UpdatedAt = time.Now()
	order.LinkPrecedence = "secondary"
	if err := DB.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}
	c.JSON(200, order)
}
