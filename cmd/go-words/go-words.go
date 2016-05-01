// Main package to build `go-words` executable.
package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"github.com/oytuntez/go-words"
)

// Allows compiling the package and running from terminal
func main() {
	run()
}

// Runs the package by getting input from terminal
// User provides file path to count the words
func run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	words := words.Words{}
	wordCount := words.File(path)

	println(wordCount)
}