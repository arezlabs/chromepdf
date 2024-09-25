const { execFile } = require('child_process');
const path = require('path');

let goBinary, chromiumBinary;

switch (process.platform) {
  case 'linux':
    goBinary = path.join(__dirname, '../go/bin/chrome-pdf_linux');
    chromiumBinary = path.join(__dirname, '../go/chromium/chrome-linux');
    break;
  case 'darwin':
    goBinary = path.join(__dirname, '../go/bin/chrome-pdf_macos');
    chromiumBinary = path.join(__dirname, '../go/chromium/chrome-macos');
    break;
  case 'win32':
    goBinary = path.join(__dirname, '../go/bin/chrome-pdf_windows.exe');
    chromiumBinary = path.join(__dirname, '../go/chromium/chrome-windows.exe');
    break;
  default:
    throw new Error('Unsupported platform');
}

function convertHTMLToPDF(htmlContent, outputFile, callback) {
  const args = [htmlContent, outputFile, chromiumBinary];
  execFile(goBinary, args, (error, stdout, stderr) => {
    if (error) {
      callback(`Error generating PDF: ${stderr}`);
    } else {
      callback(null, `PDF saved to ${outputFile}`);
    }
  });
}

module.exports = { convertHTMLToPDF };
