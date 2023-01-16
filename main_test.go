package main

import (
	"fmt"
	"testing"
)

func test(parseMode string) error {
	pCount, err := loadPrinters(parseMode)
	if err != nil {
		return fmt.Errorf("cannot load printers from CUPS: %w", err)
	}

	if pCount != cupsPrintersCount {
		return fmt.Errorf("invalid number of printers loaded from CUPS: %d instead of %d", pCount, cupsPrintersCount)
	}

	return nil
}

func TestMainDirect(t *testing.T) {
	if err := test(directParseMode); err != nil {
		t.Errorf("error = %v", err)
	}
}

func TestMainPreread(t *testing.T) {
	if err := test(prereadParseMode); err != nil {
		t.Errorf("error = %v", err)
	}
}
