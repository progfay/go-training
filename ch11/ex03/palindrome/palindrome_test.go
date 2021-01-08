package palindrome_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/progfay/go-training/ch11/ex03/palindrome"
)

var seed int64 = time.Now().UTC().UnixNano()

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r1 := rune(rng.Intn(0x0800))
		r2 := rune(rng.Intn(0x0800-1)) + 1
		runes[i] = r1
		runes[n-1-i] = r1 + r2
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !palindrome.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if palindrome.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
