package api

import (
	"context"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"dsgo/internal/client"
)

func TestPurpose(t *testing.T) {
	if Purpose("a.PDF") != "text_recognition" || Purpose("a.png") != "text_recognition" {
		t.Fatal(Purpose("a.PDF"), Purpose("a.png"))
	}
	if Purpose("notes.md") != "user_data" || Purpose("sheet.xlsx") != "user_data" {
		t.Fatal(Purpose("notes.md"), Purpose("sheet.xlsx"))
	}
}

func TestSearchFetchesBodies(t *testing.T) {
	var searchBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/tools/search":
			data, _ := io.ReadAll(r.Body)
			searchBody = string(data)
			w.Write([]byte(`{"results":[{"url":"https://example.com/a"},{"url":"https://example.com/b"}]}`))
		case "/v1/tools/fetch":
			data, _ := io.ReadAll(r.Body)
			if strings.Contains(string(data), "example.com/b") {
				http.Error(w, `{"error":{"message":"nope"}}`, http.StatusBadGateway)
				return
			}
			w.Write([]byte(`{"title":"甲","content":"正文甲"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	result, err := Search(context.Background(), client.New(srv.URL, "test"), "查询")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(searchBody, `"count":3`) || !strings.Contains(searchBody, `"need_content":false`) {
		t.Fatal(searchBody)
	}
	if len(result.Pages) != 1 || result.Pages[0].Title != "甲" || result.Pages[0].Content != "正文甲" {
		t.Fatalf("%+v", result)
	}
}

func TestChatPlainBody(t *testing.T) {
	var body string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Fatalf("path %s", r.URL.Path)
		}
		data, _ := io.ReadAll(r.Body)
		body = string(data)
		w.Write([]byte(`{"choices":[{"message":{"content":"你好"}}]}`))
	}))
	defer srv.Close()
	result, err := Chat(context.Background(), client.New(srv.URL, "test"), "deepseek-v4-pro", "问题", nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Model != "deepseek-v4-pro" || result.Content != "你好" {
		t.Fatalf("%+v", result)
	}
	if !strings.Contains(body, `"model":"deepseek-v4-pro"`) || !strings.Contains(body, `"content":"问题"`) || !strings.Contains(body, `"stream":false`) || !strings.Contains(body, `"type":"disabled"`) {
		t.Fatal(body)
	}
}

func TestChatUploadsThenAsks(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "note.txt")
	if err := os.WriteFile(path, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	var sawFile bool
	var chatBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/files":
			_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
			if err != nil {
				t.Fatal(err)
			}
			reader := multipart.NewReader(r.Body, params["boundary"])
			first, err := reader.NextPart()
			if err != nil {
				t.Fatal(err)
			}
			if first.FormName() != "purpose" {
				t.Fatalf("first field %s", first.FormName())
			}
			sawFile = true
			w.Write([]byte(`{"id":"file-api-1"}`))
		case "/chat/completions":
			data, _ := io.ReadAll(r.Body)
			chatBody = string(data)
			w.Write([]byte(`{"choices":[{"message":{"content":"好"}}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	result, err := Chat(context.Background(), client.New(srv.URL, "test"), "other-model", "总结", []string{path})
	if err != nil {
		t.Fatal(err)
	}
	if !sawFile {
		t.Fatal("missing upload")
	}
	if result.Model != "deepseek-flash" || result.Content != "好" {
		t.Fatalf("%+v", result)
	}
	if !strings.Contains(chatBody, `"file_id":"file-api-1"`) || !strings.Contains(chatBody, `"type":"disabled"`) || strings.Contains(chatBody, `"stream":true`) {
		t.Fatal(chatBody)
	}
}

func TestOCRSendsPurposeFirst(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "page.png")
	if err := os.WriteFile(path, []byte("img"), 0o600); err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			t.Fatal(err)
		}
		reader := multipart.NewReader(r.Body, params["boundary"])
		first, err := reader.NextPart()
		if err != nil {
			t.Fatal(err)
		}
		if first.FormName() != "purpose" {
			t.Fatalf("first field %s", first.FormName())
		}
		w.Write([]byte(`{"pages":[{"index":0,"text":"甲"},{"index":1,"text":"乙"}]}`))
	}))
	defer srv.Close()
	result, err := OCR(context.Background(), client.New(srv.URL, "test"), path)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Pages) != 2 || result.Pages[1].Text != "乙" {
		t.Fatalf("%+v", result)
	}
}

func TestMissingFileDoesNotCall(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("unexpected request")
	}))
	defer srv.Close()
	_, err := OCR(context.Background(), client.New(srv.URL, "test"), filepath.Join(t.TempDir(), "missing.png"))
	if err == nil || !strings.Contains(err.Error(), "文件不存在") {
		t.Fatal(err)
	}
}
