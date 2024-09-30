package services

import (
	"Feedback_service/models"
	"Feedback_service/repository"
	"errors"
	"fmt"
)

type ReviewService struct {
	UserRepo *repository.UserRepository
}

func (s *ReviewService) SubmitReview(review models.Review) (int64, error) {
	exists, err := s.UserRepo.FindById(review.ClientID)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	if !exists {
		return 0, errors.New("client not found")
	}
	stmt, err := s.UserRepo.DB.Prepare("UPDATE feedback SET rating = ?, comment = ? WHERE client_id=?")
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	res, err := stmt.Exec(review.Rating, review.Comment, review.ClientID)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, fmt.Errorf("%w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return id, nil
}
