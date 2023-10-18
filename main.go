package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sepulCOmpany/backend/internal/db"
	"github.com/sepulCOmpany/backend/internal/handlers"
)

func main() {
	database := db.NewDataBase()
	api := handlers.NewApi(database)

	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	g.POST("/register", api.Register)
	g.POST("/login", api.Login)

	g.POST("/sepulca", api.CreateSepulca)
	g.GET("/grimziks", api.GetGrimziks)
	g.PUT("/sepulca/vaccinate", api.VaccinateSepulca)
	g.PUT("/sepulca/rubber", api.RubberSepulca)

	if err := g.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
