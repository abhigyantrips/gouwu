<div align="center">
  <img src="meta/gouwu-logo.png" alt="gouwu logo" height="120">
  <h1>gouwu</h1>
  <p><em>A Go package for transforming text into adorable uwu speak! (´｡• ᵕ •｡`) ♡</em></p>
  
  [![Go Reference](https://pkg.go.dev/badge/github.com/abhigyantrips/gouwu.svg)](https://pkg.go.dev/github.com/abhigyantrips/gouwu)
  [![Go Version](https://img.shields.io/github/go-mod/go-version/abhigyantrips/gouwu)](https://golang.org/)
  [![License](https://img.shields.io/github/license/abhigyantrips/gouwu)](LICENSE)
  [![Go Report Card](https://goreportcard.com/badge/github.com/abhigyantrips/gouwu)](https://goreportcard.com/report/github.com/abhigyantrips/gouwu)
</div>

## ✨ Features

- 🎯 **Configurable transformations** - Control word, space, and exclamation modifications
- 🎲 **Deterministic results** - Uses seeded random generation for consistent output
- 🌐 **URL-safe** - Automatically excludes URLs from transformation
- 🎨 **Rich expressions** - Includes kawaii faces, actions, and exclamations
- ⚡ **High performance** - Efficient regex-based transformations
- 🧪 **Well tested** - Comprehensive test suite included

## 📦 Installation

```bash
go get github.com/abhigyantrips/gouwu
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/abhigyantrips/gouwu"
)

func main() {
    // Create a new uwuifier with default settings
    uwuifier := gouwu.New()
    
    // Transform your text!
    result := uwuifier.UwuifySentence("This package is amazing!")
    fmt.Println(result)
    // Output: This package is amazinyg! UwU
}
```

## ⚙️ Advanced Configuration

You can customize the uwuifier behavior with various options:

```go
// Create uwuifier with custom settings
uwuifier := gouwu.New(
    gouwu.WithWords(0.8),              // 80% word transformation probability
    gouwu.WithSpaces(gouwu.SpacesModifier{
        Faces:    0.05,                // 5% chance for kawaii faces
        Actions:  0.03,                // 3% chance for actions
        Stutters: 0.15,                // 15% chance for stutters
    }),
    gouwu.WithExclamations(1.2),       // 120% exclamation intensity
)

// Use a seeded random generator for deterministic results
uwuifier.SetSeed(42)

text := "Hello world! This is a test sentence."
result := uwuifier.UwuifySentence(text)
fmt.Println(result)
```

## 🎭 Available Transformations

### Word Transformations
- `r/l` → `w` (hello → hewwo)
- `R/L` → `W` (HELLO → HEWWO)
- `n([aeiou])` → `ny$1` (no → nyo)
- `N([aeiou])` → `NY$1` (NO → NYO)
- `ove` → `uv` (love → wuv)

### Space Modifiers
- **Faces**: Random kawaii emoticons `(´｡• ᵕ •｡`) ♡`, `(◕‿◕)♡`, `OwO`, `UwU`
- **Actions**: Cute actions `*blushes*`, `*giggles*`, `*hugs*`
- **Stutters**: Word repetition `h-hewwo`, `b-but`

### Exclamations
Enhanced punctuation with cute expressions and emoticons.

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## 📖 API Reference

### Types

#### `Uwuifier`
The main struct that handles text transformations.

#### `SpacesModifier`
Configuration for space-based transformations:
```go
type SpacesModifier struct {
    Faces    float64 // Probability for kawaii faces
    Actions  float64 // Probability for cute actions  
    Stutters float64 // Probability for stutters
}
```

### Methods

#### `New(options ...Option) *Uwuifier`
Creates a new Uwuifier instance with optional configuration.

#### `UwuifySentence(sentence string) string`
Transforms a sentence into uwu speak.

#### `SetSeed(seed int64)`
Sets the random seed for deterministic results.

#### `SetWordsModifier(modifier float64)`
Sets the word transformation probability (0.0-1.0).

#### `SetSpacesModifier(modifier SpacesModifier)`
Sets the space transformation probabilities.

#### `SetExclamationsModifier(modifier float64)`
Sets the exclamation enhancement intensity.

## 📄 License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## 🙏 Acknowledgments

I had no hand in this monstrosity, and this is a good place to mention that this project was almost completely AI-generated. The complete credits go to @Schostl for breaking their head over creating the [original package](https://github.com/Schotsl/Uwuifier) in TypeScript. I just ported that project into Go because I wanted to use it for a brainrot Discord bot I'm about to build (for *actually* learning Go). The [Uwuifier](https://uwuifier.com) project has its own set of acknowledgements, which I will be quoting here.

> I was inspired to write this god awful package by
> [this](https://honk.moe/tools/owo.html) site, I would've given full credit but
> it appears he got most of the Regex from the OwO Google Chrome extension made by
> leafysweetsgarden, I would include a link but the page is offline. So short
> story short: I'd like to give a big thanks to leafysweetsgarden for the Regex
> and the inspiration!

---

<div align="center">
  Made with disgust by <a href="https://github.com/abhigyantrips">@abhigyantrips</a>
</div>