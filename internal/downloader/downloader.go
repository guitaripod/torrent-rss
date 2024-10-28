package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type Downloader struct {
	client      *http.Client
	downloadDir string
}

func NewDownloader(downloadDir string) (*Downloader, error) {
	// Create download directory if it doesn't exist
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create download directory: %w", err)
	}

	return &Downloader{
		client:      &http.Client{},
		downloadDir: downloadDir,
	}, nil
}

func (d *Downloader) findDownloadLink(pageURL string) (string, error) {
	resp, err := d.client.Get(pageURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch torrent page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	var downloadLink string
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "class" && attr.Val == "dl_Btn" {
					// Found download button, now get href
					for _, attr := range node.Attr {
						if attr.Key == "href" {
							downloadLink = attr.Val
							return
						}
					}
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(doc)

	if downloadLink == "" {
		return "", fmt.Errorf("download link not found")
	}

	// Make sure we have full URL
	if !strings.HasPrefix(downloadLink, "http") {
		downloadLink = "https://www.torrentday.com/" + strings.TrimPrefix(downloadLink, "/")
	}

	return downloadLink, nil
}

func (d *Downloader) DownloadTorrent(pageURL string) error {
	downloadLink, err := d.findDownloadLink(pageURL)
	if err != nil {
		return fmt.Errorf("failed to find download link: %w", err)
	}

	// Get the torrent file
	resp, err := d.client.Get(downloadLink)
	if err != nil {
		return fmt.Errorf("failed to download torrent: %w", err)
	}
	defer resp.Body.Close()

	// Extract filename from Content-Disposition header or URL
	filename := filepath.Base(downloadLink)
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if strings.Contains(cd, "filename=") {
			filename = strings.Split(cd, "filename=")[1]
			filename = strings.Trim(filename, `"'`)
		}
	}

	// Create the file
	filepath := filepath.Join(d.downloadDir, filename)
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
