// Package gouwu provides text transformation capabilities to convert normal text
// into "uwu speak" with configurable parameters for words, spaces, and exclamations.
//
// The package allows you to uwuify any sentence or word (excluding URLs) with many
// configurable parameters while giving access to many kawaii sentences and faces.
// It uses a seeded random generator to ensure deterministic results.
//
// Basic usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/abhigyantrips/gouwu"
//	)
//
//	func main() {
//	    uwuifier := gouwu.New()
//	    result := uwuifier.UwuifySentence("This package is amazing!")
//	    fmt.Println(result)
//	    // Output: This package is amazinyg! UwU
//	}
//
// Advanced configuration:
//
//	uwuifier := gouwu.New(
//	    gouwu.WithWords(0.8),
//	    gouwu.WithSpaces(gouwu.SpacesModifier{
//	        Faces: 0.1,
//	        Actions: 0.05,
//	        Stutters: 0.15,
//	    }),
//	    gouwu.WithExclamations(0.5),
//	)
//
// The package provides three main transformation methods:
//   - UwuifyWords: Transforms individual words (r/l -> w, n+vowel -> ny+vowel, etc.)
//   - UwuifySpaces: Adds faces, actions, or stutters between words
//   - UwuifyExclamations: Replaces exclamations with more expressive variants
//   - UwuifySentence: Combines all transformations with URL preservation
package gouwu