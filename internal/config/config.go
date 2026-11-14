package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	DefaultBaseURL = "https://api.deepseek.com"
	DefaultModel   = "deepseek-flash"
)

// Settings is the resolved runtime configuration.
// Environment variables win over the file for the key, base URL, and model.
type Settings struct {
	APIKey  string
	BaseURL string
	Model   string
	Path    string
}

// Path resolves the config file location. An empty flag uses ~/.config/dsgo/config.yaml.
func Path(flag string) (string, error) {
	if strings.TrimSpace(flag) == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".config", "dsgo", "config.yaml"), nil
	}
	path := flag
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[2:])
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

// Load reads the config file and applies environment overrides.
func Load(flag string) (Settings, error) {
	path, err := Path(flag)
	if err != nil {
		return Settings{}, err
	}
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.SetDefault("base_url", DefaultBaseURL)
	v.SetDefault("model", DefaultModel)
	if err := v.ReadInConfig(); err != nil && !missing(err) {
		return Settings{}, fmt.Errorf("读取配置失败：%s", err.Error())
	}
	settings := Settings{
		APIKey:  strings.TrimSpace(v.GetString("api_key")),
		BaseURL: strings.TrimRight(strings.TrimSpace(v.GetString("base_url")), "/"),
		Model:   strings.TrimSpace(v.GetString("model")),
		Path:    path,
	}
	if settings.BaseURL == "" {
		settings.BaseURL = DefaultBaseURL
	}
	if settings.Model == "" {
		settings.Model = DefaultModel
	}
	if env := strings.TrimSpace(os.Getenv("DEEPSEEK_API_KEY")); env != "" {
		settings.APIKey = env
	}
	if env := strings.TrimSpace(os.Getenv("DEEPSEEK_BASE_URL")); env != "" {
		settings.BaseURL = strings.TrimRight(env, "/")
	}
	if env := strings.TrimSpace(os.Getenv("DEEPSEEK_MODEL")); env != "" {
		settings.Model = env
	}
	return settings, nil
}

// Update writes only the provided fields. Nil means leave the stored value.
// Environment variables are not written into the file.
func Update(flag string, apiKey, model *string) (string, error) {
	path, err := Path(flag)
	if err != nil {
		return "", err
	}
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil && !missing(err) {
		return "", fmt.Errorf("读取配置失败：%s", err.Error())
	}
	if v.GetString("base_url") == "" {
		v.Set("base_url", DefaultBaseURL)
	}
	if v.GetString("model") == "" && model == nil {
		v.Set("model", DefaultModel)
	}
	if _, ok := v.Get("api_key").(string); !ok && apiKey == nil {
		v.Set("api_key", "")
	}
	if apiKey != nil {
		v.Set("api_key", *apiKey)
	}
	if model != nil {
		v.Set("model", *model)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return "", err
	}
	if err := v.WriteConfigAs(path); err != nil {
		return "", err
	}
	if err := os.Chmod(path, 0o600); err != nil {
		return "", err
	}
	if err := os.Chmod(filepath.Dir(path), 0o700); err != nil {
		return "", err
	}
	return path, nil
}

func missing(err error) bool {
	var notFound viper.ConfigFileNotFoundError
	return errors.As(err, &notFound) || os.IsNotExist(err)
}
