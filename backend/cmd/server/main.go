// backend/cmd/server/main.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Question struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	ImageURL    string `json:"image_url"`
	CorrectSwipe string `json:"correct_swipe"` // "left" ou "right"
	Category    string `json:"category"`
}

type SwipeRequest struct {
	QuestionID int    `json:"question_id"`
	Direction  string `json:"direction"` // "left" ou "right"
}

type SwipeResponse struct {
	Correct bool   `json:"correct"`
	Message string `json:"message"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "host=localhost port=5432 user=postgres password=votrepassword dbname=quizdb sslmode=disable"
	}
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Impossible de ping la base de données:", err)
	}

	log.Println("Connexion à PostgreSQL réussie")
}

func getRandomQuestion(c *gin.Context) {
	var question Question
	
	query := `
		SELECT id, question, image_url, correct_swipe, category 
		FROM questions 
		ORDER BY RANDOM() 
		LIMIT 1
	`
	
	err := db.QueryRow(query).Scan(
		&question.ID,
		&question.Question,
		&question.ImageURL,
		&question.CorrectSwipe,
		&question.Category,
	)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de la question"})
		return
	}
	
	// Ne pas envoyer la bonne réponse au frontend
	response := gin.H{
		"id":       question.ID,
		"question": question.Question,
		"image_url": question.ImageURL,
		"category": question.Category,
	}
	
	c.JSON(http.StatusOK, response)
}

func checkSwipe(c *gin.Context) {
	var req SwipeRequest
	
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}
	
	var correctSwipe string
	query := "SELECT correct_swipe FROM questions WHERE id = $1"
	
	err := db.QueryRow(query, req.QuestionID).Scan(&correctSwipe)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question introuvable"})
		return
	}
	
	isCorrect := req.Direction == correctSwipe
	message := "Bravo! 🎉"
	if !isCorrect {
		message = "Perdu! 😢"
	}
	
	c.JSON(http.StatusOK, SwipeResponse{
		Correct: isCorrect,
		Message: message,
	})
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()

	// Configuration CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Routes API
	api := router.Group("/api")
	{
		api.GET("/question", getRandomQuestion)
		api.POST("/swipe", checkSwipe)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Serveur démarré sur http://localhost:8080")
	router.Run(":8080")
}