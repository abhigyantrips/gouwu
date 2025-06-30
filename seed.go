package gouwu

import (
	"errors"
	"math"
)

// Seed provides deterministic random number generation based on a string seed
type Seed struct {
	a, b, c, d uint32
}

// NewSeed creates a new seeded random number generator
func NewSeed(seed string) *Seed {
	s := &Seed{}
	s.initXmur3(seed)
	return s
}

// Random generates a random float64 between min and max
func (s *Seed) Random(min, max float64) (float64, error) {
	if min > max {
		return 0, errors.New("minimum value must be below maximum value")
	}
	if min == max {
		return 0, errors.New("minimum value cannot equal maximum value")
	}

	return s.denormalize(s.sfc32(), min, max), nil
}

// RandomInt generates a random integer between min and max (inclusive)
func (s *Seed) RandomInt(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("minimum value must be below maximum value")
	}
	if min == max {
		return 0, errors.New("minimum value cannot equal maximum value")
	}

	randFloat, err := s.Random(float64(min), float64(max))
	if err != nil {
		return 0, err
	}

	return int(math.Round(randFloat)), nil
}

// denormalize maps a value from [0,1] to [min,max]
func (s *Seed) denormalize(value, min, max float64) float64 {
	return value*(max-min) + min
}

// initXmur3 initializes the PRNG state using xmur3 hash algorithm
// https://github.com/bryc/code/blob/master/jshash/PRNGs.md
func (s *Seed) initXmur3(str string) {
	h := uint32(1779033703) ^ uint32(len(str))

	for i := 0; i < len(str); i++ {
		h = s.imul32(h^uint32(str[i]), 3432918353)
		h = (h << 13) | (h >> 19)
	}

	// Generate initial state by calling the hash function 4 times
	s.a = s.xmur3Hash(&h)
	s.b = s.xmur3Hash(&h)
	s.c = s.xmur3Hash(&h)
	s.d = s.xmur3Hash(&h)
}

// xmur3Hash generates a hash value and updates the state
func (s *Seed) xmur3Hash(h *uint32) uint32 {
	*h = s.imul32(*h^(*h>>16), 2246822507)
	*h = s.imul32(*h^(*h>>13), 3266489909)
	*h ^= *h >> 16
	return *h
}

// sfc32 implements the SFC32 PRNG algorithm
// https://github.com/bryc/code/blob/master/jshash/PRNGs.md
func (s *Seed) sfc32() float64 {
	t := s.a + s.b
	s.a = s.b ^ (s.b >> 9)
	s.b = s.c + (s.c << 3)
	s.c = ((s.c << 21) | (s.c >> 11))
	s.d = s.d + 1
	t = t + s.d
	s.c = s.c + t

	// Convert to float64 in range [0,1)
	return float64(t) / 4294967296.0
}

// imul32 performs 32-bit integer multiplication (equivalent to JS Math.imul)
func (s *Seed) imul32(a, b uint32) uint32 {
	return uint32(int32(a) * int32(b))
}
