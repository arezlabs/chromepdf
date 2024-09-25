package main

import (
	"fmt"
	"os"

	"github.com/arezlabs/chromepdf/pdfgen"
)

func main() {
	// Ensure correct usage
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./chrome-pdf <htmlContent> <outputFile>")
		return
	}

	// Get HTML content and output file from the command-line arguments
	htmlContent := os.Args[1] // HTML content passed as a string
	outputFile := os.Args[2]  // Output PDF file path

	// Create a new PDF generator instance
	pdfGenerator := pdfgen.NewPDFGenerator()

	// Convert the HTML to a PDF and save it to the output file
	err := pdfGenerator.ConvertHTMLToPDF(htmlContent, outputFile)
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
	} else {
		fmt.Printf("PDF successfully saved to %s\n", outputFile)
	}
}
