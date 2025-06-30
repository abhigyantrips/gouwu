package gouwu

import (
	"strings"
	"testing"
)

func TestUwuifyWords(t *testing.T) {
    uwuifier := New(WithWords(1.0))
    
    testCases := []struct {
        input    string
        expected string
    }{
        {"Stabbed", "Stabbed"},        // No transformable patterns
        {"Tonight", "Tonyight"},       // n+vowel -> ny+vowel  
        {"Through", "Thwough"},        // r -> w
        {"Struggling", "Stwuggwing"},  // r -> w, l -> w
        {"Netherlands", "Nyethewwands"}, // N+e -> Ny+e, r->w, l->w
        {"Grandpa", "Gwandpa"},        // r -> w
        {"love", "wuv"},               // ove -> uv, but l -> w happens after, so "love" -> "luv" -> "wuv"
        {"move", "muv"},               // ove -> uv
        {"remove", "wemuv"},           // r -> w, ove -> uv
    }
    
    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            result := uwuifier.UwuifyWords(tc.input)
            if result != tc.expected {
                t.Errorf("UwuifyWords(%q) = %q, want %q", tc.input, result, tc.expected)
            }
        })
    }
}

func TestUwuifyWordsConsistency(t *testing.T) {
    uwuifier := New(WithWords(0.5)) // 50% chance
    
    input := "This is a test sentence with lots of r and l letters"
    
    // Same input should produce same output (deterministic)
    result1 := uwuifier.UwuifyWords(input)
    result2 := uwuifier.UwuifyWords(input)
    
    if result1 != result2 {
        t.Errorf("UwuifyWords not deterministic:\n%q\n%q", result1, result2)
    }
}

func TestUwuifySpacesAddsElements(t *testing.T) {
    uwuifier := New(WithSpaces(SpacesModifier{
        Faces: 1.0, Actions: 0, Stutters: 0, // Only faces, 100% chance
    }))
    
    input := "hello world"
    result := uwuifier.UwuifySpaces(input)
    
    // Should have added faces somewhere
    hasFaces := false
    for _, face := range uwuifier.Faces {
        if strings.Contains(result, face) {
            hasFaces = true
            break
        }
    }
    
    if !hasFaces {
        t.Errorf("UwuifySpaces should have added faces but got: %q", result)
    }
}

func TestUwuifyExclamationsReplaces(t *testing.T) {
    uwuifier := New(WithExclamations(1.0)) // 100% chance
    
    testCases := []string{
        "Hello!",
        "What?",
        "Wow!!",
        "Really?!",
    }
    
    for _, input := range testCases {
        t.Run(input, func(t *testing.T) {
            result := uwuifier.UwuifyExclamations(input)
            
            // Should have replaced exclamation with something from our list
            hasCustomExclamation := false
            for _, excl := range uwuifier.Exclamations {
                if strings.Contains(result, excl) {
                    hasCustomExclamation = true
                    break
                }
            }
            
            if !hasCustomExclamation {
                t.Errorf("UwuifyExclamations should have replaced exclamation in %q, got %q", input, result)
            }
        })
    }
}

func TestDeterministicBehavior(t *testing.T) {
    uwuifier := New()
    
    testSentences := []string{
        "Hello world!",
        "This is a test sentence.",
        "Random text with multiple words and punctuation!",
    }
    
    for _, sentence := range testSentences {
        t.Run(sentence, func(t *testing.T) {
            result1 := uwuifier.UwuifySentence(sentence)
            result2 := uwuifier.UwuifySentence(sentence)
            
            if result1 != result2 {
                t.Errorf("UwuifySentence not deterministic for %q:\n%q\n%q", 
                    sentence, result1, result2)
            }
        })
    }
}

func TestURLPreservation(t *testing.T) {
    uwuifier := New()
    
    testCases := []string{
        "Check this out: https://www.example.com",
        "Visit https://github.com/user/repo for more info",
    }
    
    for _, input := range testCases {
        t.Run(input, func(t *testing.T) {
            result := uwuifier.UwuifySentence(input)
            
            // URLs should remain completely intact
            if strings.Contains(input, "https://www.example.com") && 
               !strings.Contains(result, "https://www.example.com") {
                t.Errorf("URL was modified: %s -> %s", input, result)
            }
            
            if strings.Contains(input, "https://github.com") && 
               !strings.Contains(result, "https://github.com") {
                t.Errorf("URL was modified: %s -> %s", input, result)
            }
        })
    }
}

func TestZeroModifiersNoChange(t *testing.T) {
    uwuifier := New(
        WithWords(0),
        WithSpaces(SpacesModifier{Faces: 0, Actions: 0, Stutters: 0}),
        WithExclamations(0),
    )
    
    input := "This should remain completely unchanged!"
    result := uwuifier.UwuifySentence(input)
    
    if result != input {
        t.Errorf("With zero modifiers, input should be unchanged:\ninput:  %q\nresult: %q", input, result)
    }
}

func TestTransformationOccurs(t *testing.T) {
    uwuifier := New(WithWords(1.0)) // Ensure word transformation happens
    
    input := "Hello world with letters to transform"
    result := uwuifier.UwuifySentence(input)
    
    // Should be different from input (has 'l' and 'r' to transform)
    if result == input {
        t.Errorf("Expected transformation but got identical result: %q", result)
    }
    
    // Should contain 'w' replacements
    if !strings.Contains(result, "w") {
        t.Errorf("Expected 'w' replacements in result: %q", result)
    }
}