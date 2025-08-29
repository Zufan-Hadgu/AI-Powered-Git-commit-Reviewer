package repository

import (
	"encoding/json"
	"fmt"

	"gitcommit/domain/entity"
	"gitcommit/dto"
	"gitcommit/infrastructure/ai"
)

type ReviewRepository struct {
	aiClient *ai.GrokClient
}

func NewReviewRepository(aiClient *ai.GrokClient) *ReviewRepository {
	return &ReviewRepository{aiClient: aiClient}
}

func (r *ReviewRepository) ReviewCommit(commit entity.Commit) (*entity.Review, error) {
	messages := []entity.Message{
		{Role: "system", Content: "You are a code review assistant. Respond in JSON format with suggestedMessage, summary, and Score."},
		{Role: "user", Content: fmt.Sprintf("Commit message:\n%s\n\nDiff:\n%s", commit.Message, commit.Diff)},
	}

	chatResp, err := r.aiClient.GenerateChat(messages, 300)
	if err != nil {
		return nil, err
	}

	var reviewDTO dto.ReviewResponse
	err = json.Unmarshal([]byte(chatResp.Content), &reviewDTO)
	if err != nil {
		// fallback in case AI didn't return valid JSON
		reviewDTO = dto.ReviewResponse{
			Feedback: chatResp.Content,
			Suggest:  commit.Message,
			Score:    0,
		}
	}

	// Map DTO → Domain
	review := &entity.Review{
		Feedback: reviewDTO.Feedback,
		Suggest:  reviewDTO.Suggest,
		Score:    reviewDTO.Score,
	}

	return review, nil
}
