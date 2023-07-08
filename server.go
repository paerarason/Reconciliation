package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/Reconciliation/controller"
)

func main() {

	r := gin.Default()
	r.POST("/order", controller.OrderCreate)
	r.POST("/identify", controller.OrderCreate)
	r.Run()
}
