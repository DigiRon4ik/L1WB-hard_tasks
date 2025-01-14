package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// DownloadPage downloads the content of a page and saves it locally.
func DownloadPage(client *http.Client, pageURL, outputDir string) error {
	resp, err := client.Get(pageURL)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %w", pageURL, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 response: %d for %s", resp.StatusCode, pageURL)
	}

	// Parse the URL.
	u, err := url.Parse(pageURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL %s: %w", pageURL, err)
	}

	// Determine file path.
	var filePath string
	if u.Path == "/" || u.Path == "" {
		// For the root page, use "index.html" in the output directory.
		filePath = filepath.Join(outputDir, "index.html")
	} else {
		// For other pages, use the last segment of the path.
		dir := filepath.Join(outputDir, filepath.Dir(u.Path))
		fileName := filepath.Base(u.Path)
		if !strings.HasSuffix(fileName, ".html") {
			fileName += ".html" // Ensure the file has an .html extension.
		}
		filePath = filepath.Join(dir, fileName)
	}

	// Create directories.
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories for %s: %w", filePath, err)
	}

	// Save the content to the file.
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write content to file %s: %w", filePath, err)
	}

	fmt.Printf("Saved: %s\n", filePath)
	return nil
}

// ExtractLinks extracts all links from an HTML document.
func ExtractLinks(baseURL string, body io.Reader) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return links, nil
			}
			return nil, tokenizer.Err()

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						// Ignore anchor links.
						// if strings.HasPrefix(link, "#") {
						// 	break
						// }
						// Resolve relative URLs to absolute.
						absLink, err := resolveURL(baseURL, link)
						if err == nil {
							links = append(links, absLink)
						}
						break
					}
				}
			}
		default:
			continue
		}
	}
}

// resolveURL resolves a relative URL against a base URL.
func resolveURL(base, relative string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	relativeURL, err := url.Parse(relative)
	if err != nil {
		return "", err
	}
	resolved := baseURL.ResolveReference(relativeURL)
	return resolved.String(), nil
}

// DownloadSite downloads a site recursively.
func DownloadSite(client *http.Client, baseURL, outputDir string, visited map[string]bool) error {
	if visited[baseURL] {
		return nil
	}
	visited[baseURL] = true

	fmt.Printf("Downloading: %s\n", baseURL)

	// Download the page.
	resp, err := client.Get(baseURL)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %w", baseURL, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Save the page locally.
	if err := DownloadPage(client, baseURL, outputDir); err != nil {
		return err
	}

	// Extract links from the page.
	links, err := ExtractLinks(baseURL, resp.Body)
	if err != nil {
		return err
	}

	// Filter and recursively download links belonging to the same host.
	base, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	for _, link := range links {
		parsedLink, err := url.Parse(link)
		if err != nil {
			continue
		}
		if parsedLink.Host == base.Host {
			if err := DownloadSite(client, link, outputDir, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// Checking for arguments.
	if len(os.Args) < 2 {
		fmt.Println("Usage: ... <url> [output_dir]")
		return
	}

	// What and where to download.
	baseURL, _ := url.Parse(os.Args[1])
	outputDir := baseURL.Host

	// If there is a second argument.
	if len(os.Args) > 2 {
		outputDir = os.Args[2]
	}

	client := &http.Client{}
	visited := make(map[string]bool)

	if err := DownloadSite(client, baseURL.String(), outputDir, visited); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Site downloaded successfully.")
}
