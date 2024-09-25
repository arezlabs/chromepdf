package pdfgen

import "testing"

func TestConvertHTMLToPDF(t *testing.T) {
	generator := NewPDFGenerator("/path/to/chromium")
	html := "<h1>Test PDF</h1><p>This is a test.</p>"

	err := generator.ConvertHTMLToPDF(html, "test_output.pdf")
	if err != nil {
		t.Fatalf("Failed to generate PDF: %v", err)
	}
}
