package unittests_test

import (
	"github.com/aaronschweig/wwi24sea-testing-example/unittests"
	"testing"
)

func TestAdd(t *testing.T) {
	// Definiere eine Tabelle mit Testfällen
	tests := []struct {
		a, b     int // Eingabewerte für die Funktion Add
		expected int // Erwartetes Ergebnis
	}{
		{a: 1, b: 2, expected: 3},
		{a: -1, b: -1, expected: -2},
		{a: 0, b: 0, expected: 0},
		{a: -5, b: 5, expected: 0},
		{a: 10, b: 20, expected: 30},
	}

	// Iteriere durch die Testfälle
	for _, tt := range tests {
		// Rufe die Add-Funktion auf und überprüfe das Ergebnis
		result := unittests.Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
