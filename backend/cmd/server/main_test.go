package main

import (
	"testing"
)

func TestPrettyJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `{"name":"John","age":30,"city":"New York"}`,
			expected: "{\n  \"age\": 30,\n  \"city\": \"New York\",\n  \"name\": \"John\"\n}",
		},
		{
			input:    `{"foo":"bar","baz":[1,2,3]}`,
			expected: "{\n  \"baz\": [\n    1,\n    2,\n    3\n  ],\n  \"foo\": \"bar\"\n}",
		},
	}

	for _, tt := range tests {
		formatted, err := prettyJSON([]byte(tt.input))
		if err != nil {
			t.Errorf("prettyJSON(%q) returned error: %v", tt.input, err)
			continue
		}

		if string(formatted) != tt.expected {
			t.Errorf("prettyJSON(%q) = %q, want %q", tt.input, formatted, tt.expected)
		}
	}
}
