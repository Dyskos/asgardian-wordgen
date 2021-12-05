package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defMin   int = 2
	defMax   int = 3
	defWords int = 15
)

func parseArgs() (int, int, int, bool) {
	var min, max, words int
	var help bool
	args := os.Args[1:]

	if len(args) >= 1 {
		if args[0] == "help" || args[0] == "--help" || args[0] == "-h" || args[0] == "-?" {
			help = true
			return min, max, words, help
		}
	}
	if len(args) == 0 {
		min, max, words = defMin, defMax, defWords
		fmt.Println("Min set to ", min)
		fmt.Println("Max set to ", max)
		fmt.Println("Words set to ", words)
	}
	if len(args) == 1 {
		min, max, words = parseOne(args)
	}
	if len(args) == 2 {
		min, max, words = parseTwo(args)
	}
	if len(args) >= 3 {
		min, max, words = parseThree(args)
	}
	fmt.Println()

	return min, max, words, help
}

// Parses One args
func parseOne(args []string) (int, int, int) {
	min := 2
	fmt.Println("Min set to ", min)

	max := 3
	fmt.Println("Max set to ", max)

	words, err := strconv.Atoi(args[0])
	if err != nil {
		words = 15
		fmt.Println("Invalid argument: Words set to ", words)
	} else if words < 1 {
		words = 15
		fmt.Println("Invalid value: Words set to ", words)
	} else {
		fmt.Println("Words set to ", words)
	}

	return min, max, words
}

// Parses Two args
func parseTwo(args []string) (int, int, int) {
	min := 2
	fmt.Println("Min set to ", min)

	max, err := strconv.Atoi(args[0])
	if err != nil {
		max = min + 1
		fmt.Println("Invalid type: Max set to ", max)
	} else if max < 2 {
		max = min + 1
		fmt.Println("Invalid value: Max set to ", max)
	} else {
		fmt.Println("Max set to ", max)
	}

	words, err := strconv.Atoi(args[1])
	if err != nil {
		words = 15
		fmt.Println("Invalid type: Words set to ", words)
	}
	if words < 1 {
		words = 15
		fmt.Println("Invalid value: Words set to ", words)
	} else {
		fmt.Println("Words set to ", words)
	}

	return min, max, words
}

// Parses Three args
func parseThree(args []string) (int, int, int) {
	min, err := strconv.Atoi(args[0])
	if err != nil {
		min = 2
		fmt.Println("Invalid type: Min set to ", min)
	} else if min < 2 {
		min = 2
		fmt.Println("Invalid value: Min set to ", min)
	} else {
		fmt.Println("Min set to ", min)
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		max = min + 1
		fmt.Println("Invalid type: Max set to ", max)
	} else if max < min {
		max = min + 1
		fmt.Println("Invalid value: Max cannot be less than Min: Max set to ", max)
	} else {
		fmt.Println("Max set to ", max)
	}

	words, err := strconv.Atoi(args[2])
	if err != nil {
		words = 15
		fmt.Println("Invalid type: Words set to ", words)
	} else if words < 1 {
		words = 15
		fmt.Println("Invalid value: Words set to ", words)
	} else {
		fmt.Println("Words set to ", words)
	}

	return min, max, words
}

// TODO: improve later
func printHelp() {
	fmt.Println("usage:")
	fmt.Println("\twordgen [help][--help][-h][-?]")
	fmt.Println("\twordgen [2...][2...][1...]")
}
