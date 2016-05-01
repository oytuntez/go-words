package base

import (
	"io/ioutil"
	"strings"
	"github.com/oytuntez/go-words/counters/registry"
	"github.com/oytuntez/go-words/counters/interface"
)

var name string = "base"

type Counter struct {
	interfaces.CounterInterface
	name string
}

func (c Counter) Count(s string) int {
	words := strings.Fields(s)

	return len(words)
}

func (c Counter) Text(s string) int {
	words := strings.Fields(s)

	return len(words)
}

func (c Counter) File(p string) int {
	dat, err := ioutil.ReadFile(p)

	if err != nil {
		panic(err)
	}

	return c.Text(string(dat))
}

func init() {
	registry.Add(name, &Counter{})
}