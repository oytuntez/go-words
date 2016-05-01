package main

import (
	"path/filepath"
	_ "github.com/oytuntez/go-words/counters/all"
	"github.com/oytuntez/go-words/counters/interface"
	"github.com/oytuntez/go-words/counters/registry"
	"bufio"
	"os"
	"fmt"
	"strings"
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

	words := Words{}
	wordCount := words.File(path)

	println(wordCount)
}

// Package struct. Provides File and Text methods.
type Words struct {}

// Detects the type of a file by extension.
// This is used in File function to match related word counters under `counters` directory (and registry)
func (w *Words) detectType(p string) string {
	if p == "" {
		panic("Please provide file path.")
	}

	// Extension includes dot: ".csv"
	var extension = filepath.Ext(p)

	if len(extension) > 0 {
		// Remove the leading dot
		extension = extension[1:len(extension)]
	} else {
		panic("No extension.")
	}

	return extension
}

// Takes a file path and returns the word count by using related word counter for this file type.
// If a counter for this file type is not available, falls back to default counter, which is `base` package.
func (w *Words) File(p string) int {
	var fileType string = w.detectType(p)
	var counter interfaces.CounterInterface = registry.Get(fileType)
	var wordCount int = counter.File(p)

	return wordCount
}

// Takes a string and returns the word count by using Txt counter.
func (w *Words) Text(s string) int {
	var counter interfaces.CounterInterface = registry.Get("txt")
	var wordCount int = counter.Text(s)

	return wordCount
}
