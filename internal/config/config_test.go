package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUpdateKeepsUnspecifiedFields(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("DEEPSEEK_API_KEY", "from-env")
	t.Setenv("DEEPSEEK_BASE_URL", "https://env.example")
	path := filepath.Join(dir, "config.yaml")
	key := "file-key"
	if _, err := Update(path, &key, nil); err != nil {
		t.Fatal(err)
	}
	model := "deepseek-v4-pro"
	written, err := Update(path, nil, &model)
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(written)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("mode %v", info.Mode().Perm())
	}
	t.Setenv("DEEPSEEK_API_KEY", "")
	t.Setenv("DEEPSEEK_BASE_URL", "")
	t.Setenv("DEEPSEEK_MODEL", "")
	settings, err := Load(written)
	if err != nil {
		t.Fatal(err)
	}
	if settings.APIKey != "file-key" || settings.Model != "deepseek-v4-pro" {
		t.Fatalf("%+v", settings)
	}
	if settings.BaseURL != DefaultBaseURL {
		t.Fatalf("base %s", settings.BaseURL)
	}
}

func TestEnvOverridesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	key := "file-key"
	if _, err := Update(path, &key, nil); err != nil {
		t.Fatal(err)
	}
	model := "deepseek-v4-pro"
	if _, err := Update(path, nil, &model); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DEEPSEEK_API_KEY", "env-key")
	t.Setenv("DEEPSEEK_BASE_URL", "https://override.example")
	t.Setenv("DEEPSEEK_MODEL", "deepseek-flash")
	settings, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if settings.APIKey != "env-key" || settings.BaseURL != "https://override.example" || settings.Model != "deepseek-flash" {
		t.Fatalf("%+v", settings)
	}
}
