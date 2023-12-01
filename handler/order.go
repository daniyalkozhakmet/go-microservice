package handler

import (
	"fmt"
	"microservice/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderRepo *repository.OrderRepo
}

func (o *Order) Create(c *gin.Context) {
	fmt.Println("Create method")
}
func (o *Order) List(c *gin.Context) {
	orders, err := o.OrderRepo.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"orders": orders,
	})
}
func (o *Order) GetByID(c *gin.Context) {
	fmt.Println("Get order by id")
}
func (o *Order) UpdateByID(c *gin.Context) {
	fmt.Println("Update by id")
}
func (o *Order) DeleteByID(c *gin.Context) {
	fmt.Println("Delete by ID")
}
