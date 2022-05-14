package main

import (
	"fmt"
	appLog "gin-gorm/restApi/log"
	"gin-gorm/restApi/model"
	grocery "gin-gorm/restApi/service"
	"log"

	zerolog "gin-gorm/restApi/log/zerolog"

	"github.com/gin-gonic/gin"
)

func initLogger() appLog.AppLogger {
	appLog.AppLog = zerolog.NewLogger("Gin-gormApp")

	return appLog.AppLog
}

func main() {
	model.Database()

	router := gin.Default()
	initLogger()

	router.GET("/groceries", grocery.GetGroceries)
	router.GET("/grocery/:id", grocery.GetGrocery)
	router.POST("/grocery", grocery.PostGrocery)
	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":10002"))
	fmt.Printf("heello")
}
