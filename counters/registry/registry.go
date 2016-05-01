package registry

import (
	"github.com/oytuntez/go-words/counters/interface"
	"fmt"
)

var Counters = map[string]interfaces.CounterInterface{}
var DefaultCounter = "base"

func Add(name string, counter interfaces.CounterInterface) bool {
	fmt.Printf("Adding new counter: %v\n", name)
	Counters[name] = counter

	return true
}

func Get(name string) interfaces.CounterInterface {
	counter, ok := Counters[name]

	if !ok {
		defaultCounter, ok := Counters[DefaultCounter]

		if !ok {
			panic("Counter not found.")
		}

		counter = defaultCounter
	}

	return counter
}