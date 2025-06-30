package gouwu

import (
	"errors"
	"regexp"
	"strings"
)

// SpacesModifier defines probabilities for space transformations
type SpacesModifier struct {
	Faces    float64 `json:"faces"`
	Actions  float64 `json:"actions"`
	Stutters float64 `json:"stutters"`
}

// UwuReplacement represents a regex replacement rule
type UwuReplacement struct {
	Pattern     *regexp.Regexp
	Replacement string
}

// Default configuration values
var (
	DefaultWords        = 0.9
	DefaultSpaces       = SpacesModifier{Faces: 0.04, Actions: 0.02, Stutters: 0.1}
	DefaultExclamations = 1.0
)

// Uwuifier transforms text into uwu speak
type Uwuifier struct {
	Faces        []string
	Exclamations []string
	Actions      []string
	uwuMap       []UwuReplacement

	wordsModifier        float64
	spacesModifier       SpacesModifier
	exclamationsModifier float64
}

// Option defines a configuration function
type Option func(*Uwuifier)

// WithWords sets the word transformation probability
func WithWords(probability float64) Option {
	return func(u *Uwuifier) {
		u.SetWordsModifier(probability)
	}
}

// WithSpaces sets the space transformation probabilities
func WithSpaces(spaces SpacesModifier) Option {
	return func(u *Uwuifier) {
		u.SetSpacesModifier(spaces)
	}
}

// WithExclamations sets the exclamation transformation probability
func WithExclamations(probability float64) Option {
	return func(u *Uwuifier) {
		u.SetExclamationsModifier(probability)
	}
}

// New creates a new Uwuifier with optional configuration
func New(opts ...Option) *Uwuifier {
	u := &Uwuifier{
		Faces: []string{
			"(・`ω´・)", ";;w;;", "OwO", "UwU", ">w<",
			"^w^", "ÚwÚ", "^-^", ":3", "x3",
		},
		Exclamations: []string{"!?", "?!!", "?!?1", "!!11", "?!?!"},
		Actions: []string{
			"*blushes*", "*whispers to self*", "*cries*", "*screams*",
			"*sweats*", "*twerks*", "*runs away*", "*screeches*",
			"*walks away*", "*sees bulge*", "*looks at you*",
			"*notices buldge*", "*starts twerking*", "*huggles tightly*",
			"*boops your nose*",
		},
		wordsModifier:        DefaultWords,
		spacesModifier:       DefaultSpaces,
		exclamationsModifier: DefaultExclamations,
	}

	// Initialize uwu replacement patterns
	u.uwuMap = []UwuReplacement{
		{regexp.MustCompile(`ove`), "uv"},          // Do this FIRST
		{regexp.MustCompile(`[rl]`), "w"},          // Lowercase r/l -> w
		{regexp.MustCompile(`[RL]`), "W"},          // Uppercase R/L -> W
		{regexp.MustCompile(`n([aeiou])`), "ny$1"}, // n + vowel -> ny + vowel
		{regexp.MustCompile(`N([aeiou])`), "Ny$1"}, // N + vowel -> Ny + vowel
		{regexp.MustCompile(`N([AEIOU])`), "NY$1"}, // N + VOWEL -> NY + VOWEL
	}

	// Apply options
	for _, opt := range opts {
		opt(u)
	}

	return u
}

// Getters
func (u *Uwuifier) WordsModifier() float64         { return u.wordsModifier }
func (u *Uwuifier) SpacesModifier() SpacesModifier { return u.spacesModifier }
func (u *Uwuifier) ExclamationsModifier() float64  { return u.exclamationsModifier }

// Setters with validation
func (u *Uwuifier) SetWordsModifier(value float64) error {
	if value < 0 || value > 1 {
		return errors.New("wordsModifier value must be between 0 and 1")
	}
	u.wordsModifier = value
	return nil
}

func (u *Uwuifier) SetSpacesModifier(value SpacesModifier) error {
	sum := value.Faces + value.Actions + value.Stutters
	if sum < 0 || sum > 1 {
		return errors.New("spacesModifier sum must be between 0 and 1")
	}
	u.spacesModifier = value
	return nil
}

func (u *Uwuifier) SetExclamationsModifier(value float64) error {
	if value < 0 || value > 1 {
		return errors.New("exclamationsModifier value must be between 0 and 1")
	}
	u.exclamationsModifier = value
	return nil
}

// UwuifyWords transforms words using regex patterns
func (u *Uwuifier) UwuifyWords(sentence string) string {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		if isAt(word) || isURI(word) {
			continue
		}

		seed := NewSeed(word)

		for _, replacement := range u.uwuMap {
			// Generate random value for each pattern
			randVal, _ := seed.Random(0, 1)
			if randVal > u.wordsModifier {
				continue
			}

			word = replacement.Pattern.ReplaceAllString(word, replacement.Replacement)
		}

		words[i] = word
	}

	return strings.Join(words, " ")
}

// UwuifySpaces transforms spaces by adding faces, actions, or stutters
func (u *Uwuifier) UwuifySpaces(sentence string) string {
	words := strings.Split(sentence, " ")

	faceThreshold := u.spacesModifier.Faces
	actionThreshold := u.spacesModifier.Actions + faceThreshold
	stutterThreshold := u.spacesModifier.Stutters + actionThreshold

	for i, word := range words {
		if word == "" {
			continue
		}

		seed := NewSeed(word)
		randVal, _ := seed.Random(0, 1)

		firstChar := string(word[0])

		checkCapital := func() {
			// Check if we should remove the first capital letter
			if firstChar != strings.ToUpper(firstChar) {
				return
			}
			// If word has higher than 50% upper case
			if getCapitalPercentage(word) > 0.5 {
				return
			}

			// If it's the first word
			if i == 0 {
				word = strings.ToLower(firstChar) + word[1:]
			} else {
				prevWord := words[i-1]
				if len(prevWord) > 0 {
					lastChar := prevWord[len(prevWord)-1]
					punctuation := regexp.MustCompile(`[.!?\-]`)
					if punctuation.MatchString(string(lastChar)) {
						word = strings.ToLower(firstChar) + word[1:]
					}
				}
			}
		}

		if randVal <= faceThreshold && len(u.Faces) > 0 && !isBreak(word) {
			// Add random face
			faceIdx, _ := seed.RandomInt(0, len(u.Faces)-1)
			word += " " + u.Faces[faceIdx]
			checkCapital()
		} else if randVal <= actionThreshold && len(u.Actions) > 0 && !isBreak(word) {
			// Add random action
			actionIdx, _ := seed.RandomInt(0, len(u.Actions)-1)
			word += " " + u.Actions[actionIdx]
			checkCapital()
		} else if randVal <= stutterThreshold && !isURI(word) && !isBreak(word) {
			// Add stutter
			stutterCount, _ := seed.RandomInt(0, 2)
			stutter := strings.Repeat(firstChar+"-", stutterCount)
			word = stutter + word
		}

		words[i] = word
	}

	return strings.Join(words, " ")
}

// UwuifyExclamations replaces exclamations with more expressive ones
func (u *Uwuifier) UwuifyExclamations(sentence string) string {
	words := strings.Split(sentence, " ")
	pattern := regexp.MustCompile(`[?!]+$`)

	for i, word := range words {
		seed := NewSeed(word)
		randVal, _ := seed.Random(0, 1)

		if !pattern.MatchString(word) ||
			randVal > u.exclamationsModifier ||
			isBreak(word) {
			continue
		}

		word = pattern.ReplaceAllString(word, "")
		exclamationIdx, _ := seed.RandomInt(0, len(u.Exclamations)-1)
		word += u.Exclamations[exclamationIdx]

		words[i] = word
	}

	return strings.Join(words, " ")
}

// UwuifySentence applies all transformations to a sentence
func (u *Uwuifier) UwuifySentence(sentence string) string {
	result := sentence
	result = u.UwuifyWords(result)
	result = u.UwuifyExclamations(result)
	result = u.UwuifySpaces(result)
	return result
}
