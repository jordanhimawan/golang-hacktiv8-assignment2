package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi8-assignment/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type controllerOrder struct {
	db *gorm.DB
}

func NewControllerOrder(db *gorm.DB) *controllerOrder {
	return &controllerOrder{
		db: db,
	}
}

func (in *controllerOrder) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	in.db.Find(&orders)

	if len(orders) == 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (in *controllerOrder) CreateOrder(c *gin.Context) {
	var order models.Order

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Create(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": order,
	})

}

func (in *controllerOrder) DeleteOrderByID(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("id")

	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "order deleted!",
	})
}

func (in *controllerOrder) UpdateOrderByID(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
	)

	id := c.Param("id")

	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newOrder)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Model(&order).Updates(newOrder).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "order updated!",
		"data": newOrder,
	})
}
