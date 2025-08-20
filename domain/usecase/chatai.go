package usecase

import "gitcommit/domain/entity"

type ChatAI interface {
	GenerateChat(messages []entity.Message, maxTokens int) (entity.ChatResponse, error)
}
