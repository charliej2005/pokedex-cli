package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "lorem ipsum",
			expected: []string{"lorem", "ipsum"},
		},
		{
			input:    "dolor     sit   amet   ",
			expected: []string{"dolor", "sit", "amet"},
		},
		{
			input:    "     gotta  catch     em  all",
			expected: []string{"gotta", "catch", "em", "all"},
		},
		{
			input:    "i wanna be the very best",
			expected: []string{"i", "wanna", "be", "the", "very", "best"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("lengths do not match | expected: %v, actual: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("words do not match | index: %v, expected: %v, actual: %v", i, expectedWord, word)
			}
		}
	}
}
