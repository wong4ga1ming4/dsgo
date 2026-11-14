package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"dsgo/internal/client"
)

// Purpose selects the Files API purpose for a chat attachment.
func Purpose(path string) string {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".pdf", ".doc", ".docx", ".ppt", ".pptx",
		".png", ".jpg", ".jpeg", ".gif", ".webp", ".bmp", ".tif", ".tiff":
		return "text_recognition"
	default:
		return "user_data"
	}
}

func existingFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s: 文件不存在", path)
		}
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%s: 不是文件", path)
	}
	return nil
}

func upload(ctx context.Context, c *client.Client, path, purpose string) (string, error) {
	if err := existingFile(path); err != nil {
		return "", err
	}
	data, err := c.PostFile(ctx, "/files", []client.Field{{
		Name:  "purpose",
		Value: purpose,
	}}, path, 300*time.Second)
	if err != nil {
		return "", fmt.Errorf("%s: %s", path, err.Error())
	}
	var payload struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return "", err
	}
	if payload.ID == "" {
		return "", fmt.Errorf("%s: 上传响应没有 file id", path)
	}
	return payload.ID, nil
}
