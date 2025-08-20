package dto

import (
	"encoding/json"
	"gitcommit/domain/entity"
)

// Convert domain messages to Grok DTO messages
func FromDomainMessages(msgs []entity.Message) []Message {
	result := make([]Message, len(msgs))
	for i, m := range msgs {
		result[i] = Message{
			Role:    m.Role,
			Content: m.Content,
		}
	}
	return result
}

// Convert GrokResponseDto into domain ChatResponse
func ToDomainResponse(res GrokResponseDto) entity.ChatResponse {
	content := ""
	if len(res.Choices) > 0 {
		content = res.Choices[0].Message.Content
	}

	raw, _ := json.Marshal(res)

	return entity.ChatResponse{
		Content: content,
		Raw:     string(raw),
	}
}
