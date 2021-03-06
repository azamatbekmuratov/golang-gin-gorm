package grocery

import (
	"errors"
	appLog "gin-gorm/restApi/log"
	"gin-gorm/restApi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type GroceryUpdate struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetGroceries(c *gin.Context) {
	var groceries []model.Grocery
	appLog.AppLog.Info("App log test")

	if err := model.DB.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func GetGrocery(c *gin.Context) {
	var grocery model.Grocery

	if err := model.DB.Where("id=?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func PostGrocery(c *gin.Context) {

	var grocery NewGrocery

	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGrocery := model.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}
	if err := model.DB.Create(&newGrocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		log.Error().Err(errors.New("error in post grocery")).Msg("")
	}
}

func UpdateGrocery(c *gin.Context) {
	var grocery model.Grocery

	if err := model.DB.Where("id=?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found"})
		return
	}

	var updateGrocery GroceryUpdate

	if err := c.ShouldBindJSON(&updateGrocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.DB.Model(&grocery).Updates(model.Grocery{Name: updateGrocery.Name, Quantity: updateGrocery.Quantity}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grocery)
}

func DeleteGrocery(c *gin.Context) {

	var grocery model.Grocery

	if err := model.DB.Where("id = ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}

	if err := model.DB.Delete(&grocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted"})

}
