package pdfgen

import (
	"context"
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
		// Set the path to the Chromium binary in the linux folder
		chromiumPath = filepath.Join(".", "chromium", "linux", "chrome-linux", "chrome")
	default:
		fmt.Println("Unsupported platform")
		os.Exit(1)
	}

	return &PDFGenerator{
		ChromiumPath: chromiumPath,
	}
}

// ConvertHTMLToPDF converts HTML content to a PDF file using the Chromium binary.
func (p *PDFGenerator) ConvertHTMLToPDF(htmlContent, outputPath string) error {
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(p.ChromiumPath),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-software-rasterizer", true),
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

	// Generate PDF from the HTML content
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

	// Write the generated PDF to the output file
	err = os.WriteFile(outputPath, pdfBuffer, 0644)
	if err != nil {
		return err
	}

	return nil
}
