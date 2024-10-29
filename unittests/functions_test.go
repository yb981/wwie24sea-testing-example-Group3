package unittests_test

import (
	"testing"

	"github.com/aaronschweig/wwi24sea-testing-example/unittests"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		name           string
		a, b           int
		expectedResult int
	}{
		{
			name:           "1 + 1 = 2",
			a:              1,
			b:              1,
			expectedResult: 2,
		},
		{
			name:           "1 + -2 = -1",
			a:              1,
			b:              -2,
			expectedResult: -1,
		},
	} {
		out := unittests.Add(test.a, test.b)

		if out != test.expectedResult {
			t.Errorf("Expected %d, got %d", test.expectedResult, out)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, test := range []struct {
		name  string
		input string
		out   string
	}{
		{
			name:  "should reverse the string",
			input: "Hallo",
			out:   "ollaH",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			reversedString := unittests.Reverse(test.input)

			if reversedString != test.out {
				t.Errorf("Expected %s, got %s", test.out, reversedString)
			}
		})
	}
}
