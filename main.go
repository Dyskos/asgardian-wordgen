// Program Wordgen generates valid Asgardian roots using a markov chain to
// determine a series of phonemes. It takes 0-3 command-line arguments to
// determine generated word parameters. It currently is able to handle invalid
// input and discards words that have already been generated at runtime.

// TODO: add the ability to fetch roots already in use from the GitHub Lexicon
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var wordList []string

	rand.Seed(time.Now().UnixNano())

	syllMin, syllMax, wordLimit, help := parseArgs()
	if help {
		printHelp()
	}

	for wordNum := 0; wordNum < wordLimit; wordNum++ {
		word := word(syllMin, syllMax)
		// bad words don't count
		if ok := verifyWord(wordList, word); !ok {
			wordNum--
		} else {
			wordList = append(wordList, word)
			fmt.Println(word)
		}
	}
}
