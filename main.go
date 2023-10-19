package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sepulCOmpany/backend/internal/db"
	"github.com/sepulCOmpany/backend/internal/handlers"
)

func main() {
	database := db.NewDataBase()
	api := handlers.NewApi(database)

	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
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
	g.PUT("/sepulca/delivery", api.ChangeDeliveryState)
	g.GET("/sepulcas", api.GetAllSepulcas)

	if err := g.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
