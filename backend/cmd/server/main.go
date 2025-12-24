// Package main Quiz Swipe API
//
// @title Quiz Swipe API
// @version 1.0
// @description API pour le jeu de quiz swipe
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	"github.com/Yannstj/roots-pwa/internal/database"
	"github.com/Yannstj/roots-pwa/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Yannstj/roots-pwa/docs"
)

func main() {
	// Initialisation de la base de données
	database.Init()
	defer database.Close()

	router := gin.Default()

	// Configuration CORS (localhost + réseau local)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://192.168.1.131:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Routes API
	api := router.Group("/api")
	{
		api.GET("/question", handlers.GetRandomQuestion)
		api.POST("/swipe", handlers.CheckSwipe)
	}

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
