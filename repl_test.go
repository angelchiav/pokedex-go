package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world    ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "Rihanna, banana,    katana ain't no mountain high enough",
			expected: []string{"Rihanna", "banana", "katana", "ain't", "no", "mountain", "high", "enough"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %d words, got %d: %v", len(c.expected), len(actual), actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("at index %d: expected %q, got %q", i, c.expected[i], actual[i])
			}
		}
	}

}
