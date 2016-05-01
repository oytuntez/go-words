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

func main() {
	run()
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	words := Words{}
	wordCount := words.File(path)

	println(wordCount)
}

type Words struct {}

func (w *Words) detectType(p string) string {
	if p == "" {
		panic("Please provide file path.")
	}

	var extension = filepath.Ext(p)

	if len(extension) > 0 {
		extension = extension[1:len(extension)]
	} else {
		panic("No extension.")
	}

	return extension
}

func (w *Words) File(p string) int {
	var fileType string = w.detectType(p)
	var counter interfaces.CounterInterface = registry.Get(fileType)
	var wordCount int = counter.File(p)

	return wordCount
}

func (w *Words) Text(s string) int {
	var counter interfaces.CounterInterface = registry.Get("txt")
	var wordCount int = counter.Text(s)

	return wordCount
}
