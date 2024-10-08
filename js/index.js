const { execFile } = require('child_process');
const path = require('path');

const goBinary = path.join(__dirname, '..', 'go', 'bin', 'chrome-pdf');

// Function to convert HTML to PDF and save it to a file
function convertHTMLToPDF(htmlContent, outputFile, callback) {
  const args = [htmlContent, outputFile];
  execFile(goBinary, args, (error, stdout, stderr) => {
    if (error) {
      callback(`Error generating PDF: ${stderr}`);
    } else {
      callback(null, `PDF saved to ${outputFile}`);
    }
  });
}

// Function to convert HTML to PDF and return Base64-encoded PDF
function convertHTMLToPDFBase64(htmlContent, callback) {
  const args = [htmlContent, '--base64'];
  execFile(goBinary, args, (error, stdout, stderr) => {
    if (error) {
      callback(`Error generating Base64 PDF: ${stderr}`);
    } else {
      callback(null, stdout.trim()); // The Base64-encoded PDF is returned in stdout
    }
  });
}

module.exports = { convertHTMLToPDF, convertHTMLToPDFBase64 };