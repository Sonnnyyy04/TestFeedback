package main

import (
	"Feedback_service/handlers"
	"Feedback_service/repository"
	"Feedback_service/services"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	storage, err := repository.New("./storage/storage.db")
	if err != nil {
		log.Fatalf("failed to init storage: %w", err)
	}
	reviewService := &services.ReviewService{UserRepo: storage}
	reviewController := &handlers.ReviewHandler{ReviewService: reviewService}
	router.Static("/static", "./static")
	router.GET("/feedback", reviewController.RenderFeedbackPage)
	router.POST("/feedback", reviewController.SubmitReview)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start: %w", err)
	}
}
