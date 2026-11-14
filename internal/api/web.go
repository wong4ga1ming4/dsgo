package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"dsgo/internal/client"
)

type Page struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SearchResult struct {
	Query string `json:"query"`
	Pages []Page `json:"pages"`
}

func Fetch(ctx context.Context, c *client.Client, rawURL string) (Page, error) {
	data, err := c.PostJSON(ctx, "/v1/tools/fetch", map[string]string{"url": rawURL}, 60*time.Second)
	if err != nil {
		return Page{}, err
	}
	var page Page
	if err := json.Unmarshal(data, &page); err != nil {
		return Page{}, err
	}
	return page, nil
}

// Search fetches the body of each hit. Failed pages are skipped.
func Search(ctx context.Context, c *client.Client, query string) (SearchResult, error) {
	data, err := c.PostJSON(ctx, "/v1/tools/search", map[string]any{
		"query":        query,
		"count":        3,
		"need_content": false,
	}, 60*time.Second)
	if err != nil {
		return SearchResult{}, err
	}
	var payload struct {
		Results []struct {
			URL string `json:"url"`
		} `json:"results"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return SearchResult{}, err
	}
	if len(payload.Results) == 0 {
		return SearchResult{}, fmt.Errorf("没有搜索结果")
	}
	pages := make([]Page, 0, len(payload.Results))
	for _, result := range payload.Results {
		if result.URL == "" {
			continue
		}
		page, err := Fetch(ctx, c, result.URL)
		if err != nil {
			continue
		}
		pages = append(pages, page)
	}
	if len(pages) == 0 {
		return SearchResult{}, fmt.Errorf("抓取搜索结果失败")
	}
	return SearchResult{Query: query, Pages: pages}, nil
}
