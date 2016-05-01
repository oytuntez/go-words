// Registry keeps track of available counters.
package registry

import (
	"github.com/oytuntez/go-words/counters/interface"
	"fmt"
)

// List of available counters.
var counters = map[string]interfaces.CounterInterface{}
// Default counter.
// (TODO) This could also be provided as configuration.
var defaultCounter = "base"

// Add a new counter to the registry.
func Add(name string, counter interfaces.CounterInterface) bool {
	counters[name] = counter

	return true
}

// Get a counter previously registered. If not available, fall back to default counter.
func Get(name string) interfaces.CounterInterface {
	counter, ok := counters[name]

	if !ok {
		defaultCounter, ok := counters[defaultCounter]

		if !ok {
			panic("Counter not found.")
		}

		counter = defaultCounter
	}

	return counter
}