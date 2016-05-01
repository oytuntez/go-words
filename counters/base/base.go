// Base counter is the default counter which also provides Count, Text and File methods.
// Most counters can share Count method to finally count the words (after processing a file in their own way)
// Other counters embed Counter struct and override its methods.
package base

import (
	"io/ioutil"
	"strings"
	"github.com/oytuntez/go-words/counters/registry"
	"github.com/oytuntez/go-words/counters/interface"
)

// Name of the counter. Used to register the counter in registry.
var name string = "base"

// Default Counter struct to embed in other counters.
type Counter struct {
	interfaces.CounterInterface
	name string
}

// Count the number of words in a string. Counters usually pass their result to this method to count the actual words.
// E.g, CSV counter would process a CSV file and strips unnecessary, non-word content,
// leaving only actual words; and then passes this naked string to this method.
func (c Counter) Count(s string) int {
	words := strings.Fields(s)

	return len(words)
}

// Takes a string and returns number of words.
// E.g, CSV counter would take a string that is actually the content of a CSV file.
// File method uses this method. The only thing File method does is to read the file and pass its content to Text method.
func (c Counter) Text(s string) int {
	words := strings.Fields(s)

	return len(words)
}

// Takes a file path, reads its content and passes the content to Text method.
func (c Counter) File(p string) int {
	dat, err := ioutil.ReadFile(p)

	if err != nil {
		panic(err)
	}

	return c.Text(string(dat))
}

// Registers this counter
func init() {
	registry.Add(name, &Counter{})
}