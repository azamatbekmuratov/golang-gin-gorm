package main

import (
	"gin-gorm/restApi/model"
	grocery "gin-gorm/restApi/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	model.Database()

	router := gin.Default()

	router.GET("/groceries", grocery.GetGroceries)
	router.GET("/grocery/:id", grocery.GetGrocery)
	router.POST("/grocery", grocery.PostGrocery)
	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":10000"))
}
