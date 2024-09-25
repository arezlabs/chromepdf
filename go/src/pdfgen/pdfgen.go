package pdfgen

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// PDFGenerator provides functionality to convert HTML to PDF using headless Chromium.
type PDFGenerator struct {
	ChromiumPath string
}

// NewPDFGenerator creates a new instance of PDFGenerator based on the OS.
func NewPDFGenerator() *PDFGenerator {
	var chromiumPath string

	// Detect platform and set the correct Chromium binary path
	switch runtime.GOOS {
	case "linux":
		// Point to the pre-downloaded Chromium binary for Linux
		chromiumPath = filepath.Join("..", "..", "chromium", "linux", "chrome-linux", "chrome")
	case "darwin":
		// For macOS, adjust the path to the Chromium binary
		chromiumPath = filepath.Join("..", "..", "chromium", "macos", "chrome-mac", "Chromium.app", "Contents", "MacOS", "Chromium")
	case "windows":
		// For Windows, adjust the path
		chromiumPath = filepath.Join("..", "..", "chromium", "windows", "chrome-win", "chrome.exe")
	default:
		fmt.Println("Unsupported platform")
		os.Exit(1)
	}

	return &PDFGenerator{
		ChromiumPath: chromiumPath,
	}
}

// ConvertHTMLToPDF converts HTML content to a PDF file.
func (p *PDFGenerator) ConvertHTMLToPDF(htmlContent, outputPath string) error {
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(p.ChromiumPath),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Create a temporary HTML file to render to PDF
	tmpFile, err := os.CreateTemp("", "temp.html")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(htmlContent))
	if err != nil {
		return err
	}

	var pdfBuffer []byte
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("file://%s", tmpFile.Name())),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuffer, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.5).
				WithPaperHeight(11).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return err
	}

	// Write the PDF to a file
	err = os.WriteFile(outputPath, pdfBuffer, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ConvertHTMLToPDFBase64 converts HTML content to a PDF and returns a Base64 encoded string.
func (p *PDFGenerator) ConvertHTMLToPDFBase64(htmlContent string) (string, error) {
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(p.ChromiumPath),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Create a temporary HTML file to render to PDF
	tmpFile, err := os.CreateTemp("", "temp.html")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(htmlContent))
	if err != nil {
		return "", err
	}

	var pdfBuffer []byte
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("file://%s", tmpFile.Name())),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuffer, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.5).
				WithPaperHeight(11).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return "", err
	}

	// Encode the PDF to Base64
	base64PDF := base64.StdEncoding.EncodeToString(pdfBuffer)

	return base64PDF, nil
}
