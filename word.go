//SUPPOSE WE HAVE A TYPEWRITER OF CONTAINING A SINGLE LINE OF CHARACTERS: a, b, c, t
//IF WE INPUT SOME WORDS, WE LIKE TO GET THE LONGEST WORD THAT CAN BE INPUT BY THE FIRST LINE.
//PRINT THE LATEST WORD, IF THEY ARE OF SAME SIZE.
//INPUT: bat, cat, rat
//OUTPUT: cat

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var allowed_chars = []rune{'a', 'b', 'c', 't'}

func difference(a []rune, b []rune) bool {

	for _, a_value := range a {
		var found int
		for _, b_value := range b {
			if a_value == b_value {
				found = 1
				break
			}
		}
		if found == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter words: ")
	reader := bufio.NewReader(os.Stdin)
	value, error := reader.ReadString('\n')
	if error != nil {
		panic(error)
	}
	value = strings.TrimSpace(strings.TrimSuffix(value, "\n"))
	words := strings.Split(value, " ")
	var output_word string
	for _, value := range words {
		var current = []rune(value)
		all_exists := difference(current[:], allowed_chars[:])
		if all_exists && len(current) >= len(output_word) {
			output_word = string(current)
		}
	}
	fmt.Println(output_word)
}
