package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct{
		input    string		
		expected []string		
	}{
		{"help", []string{"help"}},
		{"help me", []string{"help", "me"}},
		{"  help me   ", []string{"help", "me"}},
		{"", []string{}},
		{"      ", []string{}},

	}
	for _, test := range tests {
		res := cleanInput(test.input)
		if len(res) != len(test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, res)
		}
		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Expected %v, got %v", test.expected, res)
			}
		}
	}
}