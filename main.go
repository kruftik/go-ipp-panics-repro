package main

import (
	"log"
	"os"
)

const (
	cupsEndpoint      = "http://127.0.0.1:10631"
	cupsPrintersCount = 1446

	prereadParseMode = "preread"
	directParseMode  = "direct"
)

func main() {
	parseMode := os.Getenv("HTTP_RESPONSE_BODY_PARSE_MODE")
	if parseMode != directParseMode && parseMode != prereadParseMode {
		log.Printf("[ERROR] invalid parse mode, only %s and %s are allowed", directParseMode, prereadParseMode)
	}

	pCount, err := loadPrinters(parseMode)
	if err != nil {
		log.Fatalf("cannot load printers: %v", err)
		return
	}

	log.Printf("[INFO] %d printers read", pCount)
}
