package main

import (
	"math/rand"
)

type state struct {
	phonemes []string
	chances  []int
}

// a e h i k m n o p r s t u + 'start', the starting state
var chain = map[string]state{
	// the starting point in the chain, can never be returned to again
	"start": {
		[]string{"a", "e", "h", "i", "k", "m", "n", "o", "p", "r", "s", "t", "u"},
		[]int{8, 8, 5, 8, 5, 5, 5, 8, 5, 5, 5, 5, 8},
	},
	"a": {
		[]string{"h", "i", "k", "m", "n", "p", "r", "s", "t", "u"},
		[]int{1, 4, 1, 1, 1, 1, 1, 1, 1, 4},
	},
	"e": { //
		[]string{"h", "k", "m", "n", "p", "r", "s", "t", "u"},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 8},
	},
	"h": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"i": {
		[]string{"a", "h", "k", "m", "n", "o", "p", "r", "s", "t", "u"},
		[]int{8, 3, 3, 3, 3, 8, 3, 3, 3, 3, 8},
	},
	"k": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"m": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"n": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"o": {
		[]string{"h", "i", "k", "m", "n", "p", "r", "s", "t"},
		[]int{1, 8, 1, 1, 1, 1, 1, 1, 1},
	},
	"p": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"r": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"s": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"t": {
		[]string{"a", "e", "i", "o", "u"},
		[]int{1, 1, 1, 1, 1},
	},
	"u": {
		[]string{"a", "e", "h", "i", "k", "m", "n", "p", "r", "s", "t"},
		[]int{8, 8, 3, 8, 3, 3, 3, 3, 3, 3, 3},
	},
}

// function word() generates a word with the given parameters
func word(syllMin, syllMax int) string {
	word := ""
	phoneme := "start"
	syllNum := 0
	syllLimit := rand.Intn((syllMax+1)-syllMin) + syllMin
	for syllNum < syllLimit {
		phoneme = randPhoneme(chain[phoneme])
		// syllables are defined as any time a vowel appears in a root
		if phoneme == "i" ||
			phoneme == "e" ||
			phoneme == "a" ||
			phoneme == "o" ||
			phoneme == "u" {
			syllNum++
		}
		word += phoneme
	}
	return word
}

// function verifyWord() returns false if word is only 2-letter
// or has already been generated
func verifyWord(wordList []string, word string) bool {
	ok := true
	if len(word) <= 2 {
		return !ok
	} else if stringIn(wordList, word) {
		return !ok
	}
	return ok
}

// function randChain() picks a number within the range of possible next choices
func randPhoneme(phoneme state) string {
	var total, min, max int

	// determine total as a kind of "common denominator" for the integer values
	for i := range phoneme.chances {
		total = total + phoneme.chances[i]
	}

	// choose a number within the range of [1, total]
	selection := rand.Intn(total) + 1

	// interpret selection into respective letter choice
	for i := range phoneme.chances {
		if i == 0 {
			min = 0
			max = phoneme.chances[i]
		} else if i == 1 {
			min = phoneme.chances[i-1]
			max = min + phoneme.chances[i]
		} else {
			min = min + phoneme.chances[i-1]
			max = min + phoneme.chances[i]
		}
		if selection >= min+1 && selection <= max {
			selection = i
			break
		}
	}
	return phoneme.phonemes[selection]
}

// function StringInSlice checks if string is in slice, returns bool
func stringIn(sl []string, str string) bool {
	for _, v := range sl {
		if v == str {
			return true
		}
	}
	return false
}
