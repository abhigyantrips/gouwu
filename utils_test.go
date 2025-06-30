package gouwu

import (
	"fmt"
	"testing"
)

func TestIsAt(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"@username", true},
		{"@", true},
		{"username", false},
		{"", false},
		{"hello@world", false},
		{"@user123", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := isAt(tc.input)
			if result != tc.expected {
				t.Errorf("isAt(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestGetCapitalPercentage(t *testing.T) {
	testCases := []struct {
		input    string
		expected float64
	}{
		{"HELLO", 1.0},
		{"hello", 0.0},
		{"Hello", 0.2}, // 1 out of 5 letters
		{"HeLLo", 0.6}, // 3 out of 5 letters
		{"123", 0.0},   // No letters
		{"", 0.0},      // Empty string
		{"H3LL0", 1.0}, // 3 out of 3 letters (numbers don't count as letters)
		{"café", 0.0},  // Unicode support
		{"CAFÉ", 1.0},  // Unicode support
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := getCapitalPercentage(tc.input)
			if result != tc.expected {
				t.Errorf("getCapitalPercentage(%q) = %f, want %f", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsURI(t *testing.T) {
	validURIs := []string{
		"https://www.example.com",
		"http://example.com/path",
		"ftp://files.example.com",
		"mailto:user@example.com",
		"https://example.com:8080/path?query=value#fragment",
		"file:///path/to/file",
	}

	invalidURIs := []string{
		"",
		"not a uri",
		"ht tp://invalid.com",  // space in scheme
		"://example.com",       // missing scheme
		"https://exam<ple.com", // invalid character
	}

	for _, uri := range validURIs {
		t.Run("valid_"+uri, func(t *testing.T) {
			if !isURI(uri) {
				t.Errorf("isURI(%q) = false, want true", uri)
			}
		})
	}

	for _, uri := range invalidURIs {
		t.Run("invalid_"+uri, func(t *testing.T) {
			if isURI(uri) {
				t.Errorf("isURI(%q) = true, want false", uri)
			}
		})
	}
}

func TestIsBreak(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"\t\n", true},
		{"hello", false},
		{" hello ", false},
		{"\t", true},
		{"\n", true},
		{"\r\n", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input_%q", tc.input), func(t *testing.T) {
			result := isBreak(tc.input)
			if result != tc.expected {
				t.Errorf("isBreak(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestUnicodeHandling(t *testing.T) {
	// Test that our utils handle Unicode properly
	testCases := []struct {
		input string
		desc  string
	}{
		{"café", "French with accent"},
		{"こんにちは", "Japanese"},
		{"🎉🎊", "Emojis"},
		{"Москва", "Cyrillic"},
		{"العربية", "Arabic"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// These shouldn't panic or behave unexpectedly
			_ = getCapitalPercentage(tc.input)
			_ = isAt(tc.input)
			_ = isBreak(tc.input)
			_ = isURI(tc.input)
		})
	}
}

func BenchmarkIsURI(b *testing.B) {
	testURI := "https://www.example.com/path?query=value#fragment"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isURI(testURI)
	}
}

func BenchmarkGetCapitalPercentage(b *testing.B) {
	testString := "This Is A Test String With Mixed Case Letters"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getCapitalPercentage(testString)
	}
}
