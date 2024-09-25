
# ChromePDF: HTML to PDF Conversion with Headless Chromium

ChromePDF is a Go-based library that converts HTML content into a PDF using headless Chromium. It can be used as a Go package or an NPM package, allowing for easy integration in Go projects and JavaScript/Node.js environments.

## Features
- Convert HTML to PDF using headless Chromium.
- Supports both file output and Base64 encoded PDF output.
- Cross-platform support (Linux, macOS, Windows). (Only linux supported now. DEV in progress)
- Easy setup in Go and JavaScript environments (NPM).

---

## Table of Contents
1. [Installation](#installation)
   - [Go Installation](#go-installation)
   - [NPM Installation](#npm-installation)
2. [Usage](#usage)
   - [Go Usage](#go-usage)
   - [NPM Usage](#npm-usage)
3. [Examples](#examples)
4. [Contributing](#contributing)
5. [License](#license)

---

## Installation

### Go Installation

To use **ChromePDF** in your Go project, add it to your `go.mod` file:

```bash
go get github.com/arezlabs/chromepdf
```

### NPM Installation

To use **ChromePDF** in your JavaScript/Node.js project, install it using NPM:

```bash
npm install chrome-pdf
```

---

## Usage

### Go Usage

You can use ChromePDF as a Go package to convert HTML content to PDF. Below is an example of how to use it.

#### Import the Library

```go
import "github.com/arezlabs/chromepdf/pdfgen"
```

#### Example: Convert HTML to PDF and Save as File

```go
package main

import (
    "github.com/arezlabs/chromepdf/pdfgen"
    "fmt"
)

func main() {
    htmlContent := "<html><h1>Hello, PDF!</h1></html>"
    outputPath := "output.pdf"

    pdfGenerator := pdfgen.NewPDFGenerator()
    err := pdfGenerator.ConvertHTMLToPDF(htmlContent, outputPath)
    if err != nil {
        fmt.Println("Error generating PDF:", err)
    } else {
        fmt.Println("PDF created successfully:", outputPath)
    }
}
```

#### Example: Convert HTML to PDF and Get Base64 Output

```go
package main

import (
    "github.com/arezlabs/chromepdf/pdfgen"
    "fmt"
)

func main() {
    htmlContent := "<html><h1>Hello, Base64 PDF!</h1></html>"

    pdfGenerator := pdfgen.NewPDFGenerator()
    base64PDF, err := pdfGenerator.ConvertHTMLToPDFBase64(htmlContent)
    if err != nil {
        fmt.Println("Error generating Base64 PDF:", err)
    } else {
        fmt.Println("Base64 PDF:", base64PDF)
    }
}
```

### NPM Usage

You can use **ChromePDF** in your JavaScript/Node.js project by requiring the package and calling its functions.

#### Example: Convert HTML to PDF and Save as File

```javascript
const { convertHTMLToPDF } = require('chrome-pdf');

const htmlContent = '<html><h1>Hello, PDF!</h1>';
const outputFile = 'output.pdf';

convertHTMLToPDF(htmlContent, outputFile, (err, message) => {
  if (err) {
    console.error(err);
  } else {
    console.log(message);
  }
});
```

#### Example: Convert HTML to PDF and Get Base64 Output

```javascript
const { convertHTMLToPDFBase64 } = require('chrome-pdf');

const htmlContent = '<html><h1>Hello, Base64 PDF!</h1>';

convertHTMLToPDFBase64(htmlContent, (err, base64PDF) => {
  if (err) {
    console.error(err);
  } else {
    console.log('Base64 PDF:', base64PDF);
  }
});
```

---

## Examples

### Go Example (Linux)

To use this package in a Go project:

1. Install the package via `go get`:
   ```bash
   go get github.com/arezlabs/chromepdf
   ```

2. Use the following Go code to generate a PDF from HTML:
   ```go
   package main

   import (
       "github.com/arezlabs/chromepdf/pdfgen"
       "fmt"
   )

   func main() {
       htmlContent := "<html><h1>Hello, PDF!</h1></html>"
       outputPath := "output.pdf"

       pdfGenerator := pdfgen.NewPDFGenerator()
       err := pdfGenerator.ConvertHTMLToPDF(htmlContent, outputPath)
       if err != nil {
           fmt.Println("Error generating PDF:", err)
       } else {
           fmt.Println("PDF created successfully:", outputPath)
       }
   }
   ```

3. Build and run the program:
   ```bash
   go build -o chrome-pdf main.go
   ./chrome-pdf
   ```

### NPM Example

1. Install the NPM package:
   ```bash
   npm install chrome-pdf
   ```

2. Create a simple **Node.js** script (`index.js`):
   ```javascript
   const { convertHTMLToPDF } = require('chrome-pdf');

   const html = '<html><h1>Hello, PDF!</h1>';
   convertHTMLToPDF(html, 'output.pdf', (err, message) => {
     if (err) {
       console.error(err);
     } else {
       console.log(message);
     }
   });
   ```

3. Run the Node.js script:
   ```bash
   node index.js
   ```

---

## Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request. Feel free to report issues, request features, or improve the documentation.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

Let me know if you'd like to modify any sections or add more details!