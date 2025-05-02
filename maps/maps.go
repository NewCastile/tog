package main

import (
	"fmt"
	"strings"
)

var example = "How much wood would a woodchuck chuck if a woodchuck could chuck wood"

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	splittedWords := strings.Fields(s)
	for _, v := range splittedWords {
	  fmt.Println(v)
	  counter[v] = counter[v] + 1
	}
	fmt.Println(splittedWords)
	return counter
}

func main() {
	fmt.Println(WordCount(example))
}

