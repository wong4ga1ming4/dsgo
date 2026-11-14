package api

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"dsgo/internal/client"
)

type OCRPage struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

type OCRResult struct {
	Pages []OCRPage `json:"pages"`
}

func OCR(ctx context.Context, c *client.Client, path string) (OCRResult, error) {
	if err := existingFile(path); err != nil {
		return OCRResult{}, err
	}
	data, err := c.PostFile(ctx, "/v1/tools/ocr", []client.Field{{
		Name:  "purpose",
		Value: "text_recognition",
	}}, path, 300*time.Second)
	if err != nil {
		return OCRResult{}, err
	}
	var payload struct {
		Pages []OCRPage `json:"pages"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return OCRResult{}, err
	}
	if payload.Pages == nil {
		payload.Pages = []OCRPage{}
	}
	sort.SliceStable(payload.Pages, func(i, j int) bool {
		return payload.Pages[i].Index < payload.Pages[j].Index
	})
	return OCRResult{Pages: payload.Pages}, nil
}
