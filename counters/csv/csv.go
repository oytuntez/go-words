// Word counter for CSV files/content
// (TODO) CSV counter is not implemented yet.
package csv

import (
	"github.com/oytuntez/go-words/counters/base"
	"github.com/oytuntez/go-words/counters/registry"
)

var name string = "csv"

type CSV struct {
	base.Counter
}

func init() {
	registry.Add(name, &CSV{})
}