package api

import (
	"context"
	"encoding/json"
	"time"

	"dsgo/internal/client"
)

type ChatResult struct {
	Model   string `json:"model"`
	Content string `json:"content"`
}

type contentPart struct {
	Type   string `json:"type"`
	FileID string `json:"file_id,omitempty"`
	Text   string `json:"text,omitempty"`
}

// Chat sends one user turn. Files are uploaded first and referenced by id.
func Chat(ctx context.Context, c *client.Client, model, prompt string, files []string) (ChatResult, error) {
	var content any = prompt
	if len(files) > 0 {
		model = "deepseek-flash"
		parts := make([]contentPart, 0, len(files)+1)
		for _, path := range files {
			id, err := upload(ctx, c, path, Purpose(path))
			if err != nil {
				return ChatResult{}, err
			}
			parts = append(parts, contentPart{Type: "file", FileID: id})
		}
		parts = append(parts, contentPart{Type: "text", Text: prompt})
		content = parts
	}
	body := map[string]any{
		"model": model,
		"messages": []map[string]any{{
			"role":    "user",
			"content": content,
		}},
		"thinking": map[string]string{"type": "disabled"},
		"stream":   false,
	}
	timeout := 60 * time.Second
	if len(files) > 0 {
		timeout = 300 * time.Second
	}
	data, err := c.PostJSON(ctx, "/chat/completions", body, timeout)
	if err != nil {
		return ChatResult{}, err
	}
	var payload struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return ChatResult{}, err
	}
	if len(payload.Choices) == 0 {
		return ChatResult{}, errString("响应里没有回复")
	}
	return ChatResult{Model: model, Content: payload.Choices[0].Message.Content}, nil
}

type errString string

func (e errString) Error() string { return string(e) }
