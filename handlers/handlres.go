package handlers

import (
	"Feedback_service/models"
	"Feedback_service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	ReviewService *services.ReviewService
}

func (c *ReviewHandler) SubmitReview(ctx *gin.Context) {
	var review models.Review
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	_, err := c.ReviewService.SubmitReview(review)
	if err != nil {
		if err.Error() == "client not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Ссылка на голосование недоступна, свяжитесь с нами по телефону"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit review"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Thank you for your feedback"})
}

func (c *ReviewHandler) RenderFeedbackPage(ctx *gin.Context) {
	clientIDParam := ctx.Query("client_id")
	if clientIDParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}
	clientID, err := strconv.Atoi(clientIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid client_id"})
		return
	}
	exists, err := c.ReviewService.UserRepo.FindById(clientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if !exists {
		ctx.String(http.StatusNotFound, "Ссылка на голосование недоступна, свяжитесь с нами по телефону")
		return
	}
	ctx.File("index.html")
}
