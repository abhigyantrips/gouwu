package gouwu

import (
	"fmt"
	"testing"
)

func TestSeedBasicFunctionality(t *testing.T) {
	seed := NewSeed("test")

	// Test Random
	val, err := seed.Random(0, 1)
	if err != nil {
		t.Fatalf("Random() returned error: %v", err)
	}
	if val < 0 || val > 1 {
		t.Errorf("Random(0, 1) = %f, want value between 0 and 1", val)
	}

	// Test RandomInt
	intVal, err := seed.RandomInt(1, 10)
	if err != nil {
		t.Fatalf("RandomInt() returned error: %v", err)
	}
	if intVal < 1 || intVal > 10 {
		t.Errorf("RandomInt(1, 10) = %d, want value between 1 and 10", intVal)
	}
}

func TestSeedDeterminism(t *testing.T) {
	testSeeds := []string{"test", "hello", "uwu", ""}

	for _, seedStr := range testSeeds {
		t.Run(seedStr, func(t *testing.T) {
			seed1 := NewSeed(seedStr)
			seed2 := NewSeed(seedStr)

			// Test multiple calls for consistency
			for i := 0; i < 10; i++ {
				val1, err1 := seed1.Random(0, 1)
				val2, err2 := seed2.Random(0, 1)

				if err1 != nil || err2 != nil {
					t.Fatalf("Random() returned errors: %v, %v", err1, err2)
				}

				if val1 != val2 {
					t.Errorf("Call %d: seed consistency failed: %f != %f", i, val1, val2)
				}
			}
		})
	}
}

func TestSeedDifferentSeeds(t *testing.T) {
	seed1 := NewSeed("seed1")
	seed2 := NewSeed("seed2")

	val1, _ := seed1.Random(0, 1)
	val2, _ := seed2.Random(0, 1)

	// Different seeds should (very likely) produce different values
	if val1 == val2 {
		t.Logf("Warning: Different seeds produced same value: %f", val1)
		// Don't fail the test as this could theoretically happen
	}
}

func TestRandomValidation(t *testing.T) {
	seed := NewSeed("test")

	testCases := []struct {
		min, max float64
		wantErr  bool
	}{
		{0, 1, false},
		{-5, 5, false},
		{1, 0, true}, // min > max
		{5, 5, true}, // min == max
		{-10, -5, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Random(%f,%f)", tc.min, tc.max), func(t *testing.T) {
			_, err := seed.Random(tc.min, tc.max)
			if tc.wantErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestRandomIntValidation(t *testing.T) {
	seed := NewSeed("test")

	testCases := []struct {
		min, max int
		wantErr  bool
	}{
		{0, 10, false},
		{-5, 5, false},
		{10, 0, true}, // min > max
		{5, 5, true},  // min == max
		{-10, -5, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("RandomInt(%d,%d)", tc.min, tc.max), func(t *testing.T) {
			_, err := seed.RandomInt(tc.min, tc.max)
			if tc.wantErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestRandomDistribution(t *testing.T) {
	seed := NewSeed("distribution_test")
	const iterations = 10000
	const buckets = 10
	counts := make([]int, buckets)

	for i := 0; i < iterations; i++ {
		val, _ := seed.Random(0, 1)
		bucket := int(val * float64(buckets))
		if bucket >= buckets {
			bucket = buckets - 1 // Handle edge case where val == 1.0
		}
		counts[bucket]++
	}

	// Check if distribution is roughly uniform (very loose test)
	expected := iterations / buckets
	tolerance := expected / 2 // 50% tolerance

	for i, count := range counts {
		if count < expected-tolerance || count > expected+tolerance {
			t.Logf("Bucket %d: got %d, expected ~%d", i, count, expected)
		}
	}
}

func TestRandomIntRange(t *testing.T) {
	seed := NewSeed("range_test")

	testCases := []struct {
		min, max int
	}{
		{0, 9},
		{-5, 5},
		{100, 200},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Range_%d_%d", tc.min, tc.max), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				val, err := seed.RandomInt(tc.min, tc.max)
				if err != nil {
					t.Fatalf("RandomInt(%d, %d) returned error: %v", tc.min, tc.max, err)
				}
				if val < tc.min || val > tc.max {
					t.Errorf("RandomInt(%d, %d) = %d, out of range", tc.min, tc.max, val)
				}
			}
		})
	}
}

func TestSeedStateProgression(t *testing.T) {
	seed := NewSeed("progression_test")

	// Generate several values and ensure they're not all the same
	var values []float64
	for i := 0; i < 10; i++ {
		val, _ := seed.Random(0, 1)
		values = append(values, val)
	}

	// Check that we don't get the same value repeatedly
	allSame := true
	for i := 1; i < len(values); i++ {
		if values[i] != values[0] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Error("PRNG appears to be stuck - all values identical")
	}
}

func BenchmarkSeedRandom(b *testing.B) {
	seed := NewSeed("benchmark")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seed.Random(0, 1)
	}
}

func BenchmarkSeedRandomInt(b *testing.B) {
	seed := NewSeed("benchmark")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seed.RandomInt(0, 100)
	}
}

func BenchmarkSeedCreation(b *testing.B) {
	seeds := []string{"test1", "test2", "test3", "longer_seed_string"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewSeed(seeds[i%len(seeds)])
	}
}
