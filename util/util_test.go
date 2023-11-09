package util

import (
	"log/slog"
	"testing"
)

func TestGetLogger(t *testing.T) {
	GetLogger("test")
	slog.Info("print")
}

func TestCalculateTotalPages(t *testing.T) {
	result := CalculateTotalPages(1009, 100)
	if result != 11 {
		t.Fatalf("incorrect total page number, wanted result: 11, got result: %d", result)
	}
}
