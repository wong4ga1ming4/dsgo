package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"dsgo/internal/api"
)

func TestHelpAndRemovedCommands(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := execute(nil, &stdout, &stderr); code != 0 {
		t.Fatalf("help code=%d stderr=%s", code, stderr.String())
	}
	help := stdout.String()
	for _, name := range []string{"status", "config", "chat", "search", "fetch", "ocr"} {
		if !strings.Contains(help, name) {
			t.Fatalf("help missing %s\n%s", name, help)
		}
	}
	for _, name := range []string{"asr", "models", "balance", "files"} {
		stdout.Reset()
		stderr.Reset()
		if code := execute([]string{name}, &stdout, &stderr); code != 2 {
			t.Fatalf("%s code=%d stderr=%s", name, code, stderr.String())
		}
	}
}

func TestUsageAndMissingKey(t *testing.T) {
	t.Setenv("DEEPSEEK_API_KEY", "")
	t.Setenv("DEEPSEEK_BASE_URL", "")
	var stdout, stderr bytes.Buffer
	if code := execute([]string{"config"}, &stdout, &stderr); code != 2 {
		t.Fatalf("config code=%d stderr=%s", code, stderr.String())
	}
	if code := execute([]string{"search"}, &stdout, &stderr); code != 2 {
		t.Fatalf("search code=%d stderr=%s", code, stderr.String())
	}
	if code := execute([]string{"chat", "   "}, &stdout, &stderr); code != 2 {
		t.Fatalf("chat code=%d stderr=%s", code, stderr.String())
	}
	stderr.Reset()
	missingCfg := filepath.Join(t.TempDir(), "missing.yaml")
	if code := execute([]string{"--config", missingCfg, "status"}, &stdout, &stderr); code != 1 || !strings.Contains(stderr.String(), "未设置 DEEPSEEK_API_KEY") {
		t.Fatalf("status code=%d stderr=%s", code, stderr.String())
	}
}

func TestStatusWaitsForBothCalls(t *testing.T) {
	t.Setenv("DEEPSEEK_API_KEY", "test-key")
	failBalance := true
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("authorization %q", r.Header.Get("Authorization"))
		}
		switch r.URL.Path {
		case "/models":
			w.Write([]byte(`{"data":[{"id":"deepseek-flash","name":"DeepSeek-V4.1-Flash"}]}`))
		case "/user/balance":
			if failBalance {
				http.Error(w, `{"error":{"message":"余额失败"}}`, http.StatusInternalServerError)
				return
			}
			w.Write([]byte(`{"is_available":false,"balance_infos":[{"currency":"CNY","total_balance":"110.00","granted_balance":"10.00","topped_up_balance":"100.00"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	t.Setenv("DEEPSEEK_BASE_URL", srv.URL)
	t.Setenv("DEEPSEEK_MODEL", "")
	emptyCfg := filepath.Join(t.TempDir(), "empty.yaml")

	var stdout, stderr bytes.Buffer
	if code := execute([]string{"--config", emptyCfg, "status"}, &stdout, &stderr); code != 1 || stdout.Len() != 0 || !strings.Contains(stderr.String(), "余额失败") {
		t.Fatalf("fail code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
	failBalance = false
	stdout.Reset()
	stderr.Reset()
	code := execute([]string{"--config", emptyCfg, "--json", "status"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("ok code=%d stderr=%s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), `"ok": false`) || !strings.Contains(stdout.String(), `"topped_up": "100.00"`) || !strings.Contains(stdout.String(), `"default_model": "deepseek-flash"`) {
		t.Fatalf("json %s", stdout.String())
	}
	stdout.Reset()
	if code := execute([]string{"--config", emptyCfg, "status"}, &stdout, &stderr); code != 0 {
		t.Fatal(stderr.String())
	}
	want := "密钥不可用\n默认模型 deepseek-flash\ndeepseek-flash    DeepSeek-V4.1-Flash\nCNY  总额 110.00  赠送 10.00  充值 100.00\n"
	if stdout.String() != want {
		t.Fatalf("text %q", stdout.String())
	}
	cfg := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(cfg, []byte("model: deepseek-v4-pro\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	stdout.Reset()
	if code := execute([]string{"--config", cfg, "status"}, &stdout, &stderr); code != 0 || !strings.Contains(stdout.String(), "默认模型 deepseek-v4-pro\n") {
		t.Fatalf("config model code=%d stdout=%q stderr=%s", code, stdout.String(), stderr.String())
	}
	t.Setenv("DEEPSEEK_MODEL", "deepseek-flash")
	stdout.Reset()
	if code := execute([]string{"--config", cfg, "status"}, &stdout, &stderr); code != 0 || !strings.Contains(stdout.String(), "默认模型 deepseek-flash\n") {
		t.Fatalf("env model code=%d stdout=%q stderr=%s", code, stdout.String(), stderr.String())
	}
}

func TestChatTextAppendsCitations(t *testing.T) {
	text := chatText(api.ChatResult{Content: "约 39 万", Citations: []string{"https://example.com/a"}})
	if text != "约 39 万\n\nhttps://example.com/a\n" {
		t.Fatalf("%q", text)
	}
	if chatText(api.ChatResult{Content: "收到"}) != "收到\n" {
		t.Fatal(chatText(api.ChatResult{Content: "收到"}))
	}
}

func TestConfigWriteDoesNotEchoKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	var stdout, stderr bytes.Buffer
	code := execute([]string{"--config", path, "config", "--apikey", "secret-key", "--model", "deepseek-flash"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("code=%d stderr=%s", code, stderr.String())
	}
	if strings.Contains(stdout.String(), "secret-key") || !strings.Contains(stdout.String(), "已写入") {
		t.Fatalf("stdout %q", stdout.String())
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "secret-key") || !strings.Contains(string(data), "deepseek-flash") {
		t.Fatalf("file %s", data)
	}
	stdout.Reset()
	code = execute([]string{"--config", path, "--json", "config", "--model", "deepseek-v4-pro"}, &stdout, &stderr)
	if code != 0 || strings.Contains(stdout.String(), "secret-key") || !strings.Contains(stdout.String(), `"path"`) {
		t.Fatalf("json code=%d stdout=%s stderr=%s", code, stdout.String(), stderr.String())
	}
	data, err = os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "secret-key") || !strings.Contains(string(data), "deepseek-v4-pro") {
		t.Fatalf("merged %s", data)
	}
}
