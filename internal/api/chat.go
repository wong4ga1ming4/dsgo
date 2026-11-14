package api

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"dsgo/internal/client"
)

type ChatResult struct {
	Model     string   `json:"model"`
	Content   string   `json:"content"`
	Citations []string `json:"citations,omitempty"`
}

// ChatOptions selects server-side web search, thinking, and local files.
type ChatOptions struct {
	Search    bool
	DeepThink bool
	Files     []string
}

type contentPart struct {
	Type   string `json:"type"`
	FileID string `json:"file_id,omitempty"`
	Text   string `json:"text,omitempty"`
}

// Chat sends one user turn. Files are uploaded first and referenced by id.
// Search is a server-side tool. The client does not run it and does not force tool_choice.
func Chat(ctx context.Context, c *client.Client, model, prompt string, opt ChatOptions) (ChatResult, error) {
	var content any = prompt
	if len(opt.Files) > 0 || opt.Search {
		model = "deepseek-flash"
	}
	if len(opt.Files) > 0 {
		parts := make([]contentPart, 0, len(opt.Files)+1)
		for _, path := range opt.Files {
			id, err := upload(ctx, c, path, Purpose(path))
			if err != nil {
				return ChatResult{}, err
			}
			parts = append(parts, contentPart{Type: "file", FileID: id})
		}
		parts = append(parts, contentPart{Type: "text", Text: prompt})
		content = parts
	}
	thinking := "disabled"
	if opt.DeepThink {
		thinking = "enabled"
	}
	body := map[string]any{
		"model": model,
		"messages": []map[string]any{{
			"role":    "user",
			"content": content,
		}},
		"thinking": map[string]string{"type": thinking},
		"stream":   false,
	}
	if opt.Search {
		body["tools"] = []map[string]string{{"type": "deepseek_web_search"}}
	}
	timeout := 60 * time.Second
	if len(opt.Files) > 0 || opt.Search || opt.DeepThink {
		timeout = 300 * time.Second
	}
	data, err := c.PostJSON(ctx, "/chat/completions", body, timeout)
	if err != nil {
		return ChatResult{}, err
	}
	var payload struct {
		Choices []struct {
			Message struct {
				Content     string           `json:"content"`
				Annotations []chatAnnotation `json:"annotations"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return ChatResult{}, err
	}
	if len(payload.Choices) == 0 {
		return ChatResult{}, errString("响应里没有回复")
	}
	message := payload.Choices[0].Message
	return ChatResult{Model: model, Content: message.Content, Citations: citationURLs(message.Annotations)}, nil
}

type chatAnnotation struct {
	Type        string `json:"type"`
	URLCitation struct {
		URL string `json:"url"`
	} `json:"url_citation"`
}

func citationURLs(annotations []chatAnnotation) []string {
	seen := map[string]struct{}{}
	var urls []string
	for _, ann := range annotations {
		if ann.Type != "url_citation" {
			continue
		}
		url := strings.TrimSpace(ann.URLCitation.URL)
		if url == "" {
			continue
		}
		if _, ok := seen[url]; ok {
			continue
		}
		seen[url] = struct{}{}
		urls = append(urls, url)
	}
	return urls
}

type errString string

func (e errString) Error() string { return string(e) }
