package models

// Question représente une question du quiz
type Question struct {
	ID           int    `json:"id" example:"1"`
	Question     string `json:"question" example:"Poule ou Coq ?"`
	ImageURL     string `json:"image_url" example:"https://images.unsplash.com/photo-1548550023-2bdb3c5beed7"`
	CorrectSwipe string `json:"-"` // Non exposé au frontend
	Category     string `json:"category" example:"animaux"`
}

// QuestionResponse est la réponse envoyée au frontend (sans correct_swipe)
type QuestionResponse struct {
	ID       int    `json:"id" example:"1"`
	Question string `json:"question" example:"Poule ou Coq ?"`
	ImageURL string `json:"image_url" example:"https://images.unsplash.com/photo-1548550023-2bdb3c5beed7"`
	Category string `json:"category" example:"animaux"`
}

// SwipeRequest représente la requête de swipe
type SwipeRequest struct {
	QuestionID int    `json:"question_id" example:"1" binding:"required"`
	Direction  string `json:"direction" example:"right" binding:"required,oneof=left right"`
}

// SwipeResponse représente la réponse après un swipe
type SwipeResponse struct {
	Correct bool   `json:"correct" example:"true"`
	Message string `json:"message" example:"Bravo! 🎉"`
}

// HealthResponse représente la réponse du health check
type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// ErrorResponse représente une erreur API
type ErrorResponse struct {
	Error string `json:"error" example:"Erreur lors de la récupération"`
}
