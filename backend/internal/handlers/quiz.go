package handlers

import (
	"net/http"

	"github.com/Yannstj/roots-pwa/internal/database"
	"github.com/Yannstj/roots-pwa/internal/models"
	"github.com/gin-gonic/gin"
)

// GetRandomQuestion récupère une question aléatoire
// @Summary Récupérer une question aléatoire
// @Description Retourne une question aléatoire du quiz sans la réponse correcte
// @Tags quiz
// @Produce json
// @Success 200 {object} models.QuestionResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/question [get]
func GetRandomQuestion(c *gin.Context) {
	var question models.Question

	query := `
		SELECT id, question, image_url, correct_swipe, category
		FROM questions
		ORDER BY RANDOM()
		LIMIT 1
	`

	err := database.DB.QueryRow(query).Scan(
		&question.ID,
		&question.Question,
		&question.ImageURL,
		&question.CorrectSwipe,
		&question.Category,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Erreur lors de la récupération de la question",
		})
		return
	}

	// Réponse sans correct_swipe
	response := models.QuestionResponse{
		ID:       question.ID,
		Question: question.Question,
		ImageURL: question.ImageURL,
		Category: question.Category,
	}

	c.JSON(http.StatusOK, response)
}

// CheckSwipe vérifie si le swipe est correct
// @Summary Vérifier un swipe
// @Description Vérifie si la direction du swipe correspond à la bonne réponse
// @Tags quiz
// @Accept json
// @Produce json
// @Param request body models.SwipeRequest true "Swipe request"
// @Success 200 {object} models.SwipeResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/swipe [post]
func CheckSwipe(c *gin.Context) {
	var req models.SwipeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Requête invalide",
		})
		return
	}

	var correctSwipe string
	query := "SELECT correct_swipe FROM questions WHERE id = $1"

	err := database.DB.QueryRow(query, req.QuestionID).Scan(&correctSwipe)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Question introuvable",
		})
		return
	}

	isCorrect := req.Direction == correctSwipe
	message := "Bravo! 🎉"
	if !isCorrect {
		message = "Perdu! 😢"
	}

	c.JSON(http.StatusOK, models.SwipeResponse{
		Correct: isCorrect,
		Message: message,
	})
}

// HealthCheck vérifie l'état du serveur
// @Summary Health check
// @Description Vérifie que le serveur est en ligne
// @Tags health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthResponse{
		Status: "ok",
	})
}
