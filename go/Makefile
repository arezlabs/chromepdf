.PHONY: all linux macos windows clean

all: linux macos windows

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/chrome-pdf_linux main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/chrome-pdf_macos main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/chrome-pdf_windows.exe main.go

clean:
	rm -rf bin/*
