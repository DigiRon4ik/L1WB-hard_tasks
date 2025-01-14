package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Mock HTTP server for testing.
func mockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`<html><body><a href="/page1">Page 1</a><a href="#anchor">Anchor</a></body></html>`))
	})
	handler.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`<html><body><a href="/">Home</a></body></html>`))
	})
	return httptest.NewServer(handler)
}

func TestDownloadPage(t *testing.T) {
	server := mockServer()
	defer server.Close()

	outputDir := "test_output"
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(outputDir)

	client := &http.Client{}
	err := DownloadPage(client, server.URL, outputDir)
	if err != nil {
		t.Fatalf("DownloadPage failed: %v", err)
	}

	// Verify the downloaded file exists.
	if _, err := os.Stat(outputDir + "/index.html"); os.IsNotExist(err) {
		t.Fatalf("Downloaded file not found: %v", err)
	}
}

func TestExtractLinks(t *testing.T) {
	htmlData := `<html><body><a href="/page1">Page 1</a><a href="#anchor">Anchor</a></body></html>`
	expectedLinks := []string{"http://example.com/page1"}

	links, err := ExtractLinks("http://example.com", bytes.NewReader([]byte(htmlData)))
	if err != nil {
		t.Fatalf("ExtractLinks failed: %v", err)
	}

	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("Expected link %s, got %s", expectedLinks[i], link)
		}
	}
}

func TestResolveURL(t *testing.T) {
	baseURL := "http://example.com"
	tests := []struct {
		relative string
		expected string
	}{
		{"/path", "http://example.com/path"},
		{"/#anchor", "http://example.com/#anchor"},
		{"../relative", "http://example.com/relative"},
	}

	for _, test := range tests {
		result, err := resolveURL(baseURL, test.relative)
		if err != nil {
			t.Fatalf("resolveURL failed: %v", err)
		}
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestDownloadSite(t *testing.T) {
	server := mockServer()
	defer server.Close()

	outputDir := "test_output_site"
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(outputDir)

	client := &http.Client{}
	visited := make(map[string]bool)

	err := DownloadSite(client, server.URL, outputDir, visited)
	if err != nil {
		t.Fatalf("DownloadSite failed: %v", err)
	}

	// Verify files exist.
	if _, err := os.Stat(outputDir + "/index.html"); os.IsNotExist(err) {
		t.Fatalf("File not found: %v", err)
	}
	if _, err := os.Stat(outputDir + "/page1.html"); os.IsNotExist(err) {
		t.Fatalf("File not found: %v", err)
	}
}
