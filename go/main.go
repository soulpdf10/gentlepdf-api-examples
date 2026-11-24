package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var (
	apiKey  = os.Getenv("GENTLEPDF_API_KEY")
	baseURL = "https://api.gentlepdf.com"
)

func init() {
	if apiKey == "" {
		apiKey = "YOUR_API_KEY"
	}
	if envURL := os.Getenv("GENTLEPDF_BASE_URL"); envURL != "" {
		baseURL = envURL
	}
}

// MergePDFs sends multiple PDF files to be merged into a single document.
func MergePDFs(files []string, outputPath string) error {
	fmt.Printf("[*] Merging %d files with Go...\n", len(files))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		part, err := writer.CreateFormFile("files", filepath.Base(file))
		if err != nil {
			return err
		}
		if _, err = io.Copy(part, f); err != nil {
			return err
		}
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", baseURL+"/merge-pdf", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("[+] Merged PDF saved to %s (%d bytes)\n", outputPath, n)
	return nil
}

// CompressPDF sends a PDF with a quality preset to be compressed.
func CompressPDF(filePath string, quality string, outputPath string) error {
	fmt.Printf("[*] Compressing %s (quality: %s)...\n", filePath, quality)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return err
	}
	if _, err = io.Copy(part, f); err != nil {
		return err
	}

	if err := writer.WriteField("quality", quality); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", baseURL+"/compress-pdf", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("[+] Compressed PDF saved to %s (%d bytes)\n", outputPath, n)
	return nil
}

func main() {
	fmt.Println("GentlePDF Go API Client Examples")
	fmt.Println("Usage: Call MergePDFs() or CompressPDF() with your PDF documents.")
}
