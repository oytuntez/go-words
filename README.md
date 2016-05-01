# go-words
Word counter for various file types and text in Go

## Installation
To install `words.go`, simply run:
```
$ go get github.com/oytuntez/go-words
```

Make sure your PATH includes to the $GOPATH/bin directory so your commands can be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

## Usage
An example can be found under in the buildable code under `cmd`.

`$ cat cmd/go-words/go-words.go`


```go
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
```

## Building
```
$ cd cmd/go-words
$ go build
$ ./go-words
```

## Contributing with new counters for other file types
Take a look at `txt` package under counters/txt.
- Create a directory under `counters`.
- Embed `base.Counter` struct to your new counter Struct. It must implement `interfaces.CounterInterface`.
- Create `func init()` and register your new counter via `registry` package.
