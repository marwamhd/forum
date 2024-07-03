package main

import (
	"testing"

	"forum/Handlers"
)

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<p>hello</p>", "hello"},
		{"hello & <script>alert('XSS')</script>", "hello &amp; alert(&#39;XSS&#39;)"},
		{"hello\r\nworld", "hello<br>world"},
		{"    trim spaces    ", "trim spaces"},
		{"<b>bold</b> & <i>italic</i>", "bold &amp; italic"},
	}

	for _, test := range tests {
		output := Handlers.SanitizeInput(test.input)
		if output != test.expected {
			t.Errorf("SanitizeInput(%q) = %q; expected %q", test.input, output, test.expected)
		}
	}
}

func TestRemoveHTMLTags(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<p>hello</p>", "hello"},
		{"<div>text</div>", "text"},
		{"<br>", ""},
		{"<a href='link'>link</a>", "link"},
		{"<h1>Header</h1>", "Header"},
	}

	for _, test := range tests {
		output := Handlers.RemoveHTMLTags(test.input)
		if output != test.expected {
			t.Errorf("RemoveHTMLTags(%q) = %q; expected %q", test.input, output, test.expected)
		}
	}
}
