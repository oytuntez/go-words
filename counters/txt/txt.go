// Word counter for TXT files/content (or simply plain text content)
package txt

import (
	"github.com/oytuntez/go-words/counters/base"
	"github.com/oytuntez/go-words/counters/registry"
)

var name string = "txt"

type Txt struct {
	base.Counter
}

func init() {
	registry.Add(name, &Txt{})
}