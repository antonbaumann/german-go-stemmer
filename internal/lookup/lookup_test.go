package lookup

import "testing"

func TestIsVowelPositive(t *testing.T) {
	vowels := "aeiouyAOV"
	for _, v := range vowels {
		if !IsVowel(uint8(v)) {
			t.Errorf("`%c` should be a vowel", v)
		}
	}
}

func TestIsVowelNegative(t *testing.T) {
	consonants := "bcdfghjklmnpqrstvwxz"
	for _, v := range consonants {
		if IsVowel(uint8(v)) {
			t.Errorf("`%c` should not be a vowel", v)
		}
	}
}

func TestIsStopWordPositive(t *testing.T) {
	stopWords := []string {
		"damit", "dann", "der", "derselbe", "manche",
	}
	for _, word := range stopWords {
		if !IsStopWord(word) {
			t.Errorf("`%s` should be a stop word", word)
		}
	}
}

func TestIsStopWordNegative(t *testing.T) {
	stopWords := []string {
		"weithin", "wachsam", "umfassender",
	}
	for _, word := range stopWords {
		if IsStopWord(word) {
			t.Errorf("`%s` should not be a stop word", word)
		}
	}
}

func TestIsSEndingPositive(t *testing.T) {
	sEndings := "bdfghklmnrt"
	for _, v := range sEndings {
		if !IsSEnding(uint8(v)) {
			t.Errorf("`%c` should be a s-ending", v)
		}
	}
}

func TestIsSEndingNegative(t *testing.T) {
	sEndings := "aceijopqsuvwxyz"
	for _, v := range sEndings {
		if IsSEnding(uint8(v)) {
			t.Errorf("`%c` should not be a s-ending", v)
		}
	}
}

func TestIsStEndingPositive(t *testing.T) {
	sEndings := "bdfghklmnt"
	for _, v := range sEndings {
		if !IsStEnding(uint8(v)) {
			t.Errorf("`%c` should be a st-ending", v)
		}
	}
}

func TestIsStEndingNegative(t *testing.T) {
	sEndings := "aceijopqrsuvwxyz"
	for _, v := range sEndings {
		if IsStEnding(uint8(v)) {
			t.Errorf("`%c` should not be a s-ending", v)
		}
	}
}