// Loads all counters and runs their init() function.
// Every counter must be listed here. init() function of counters registers itself in counter registry.
package all

import (
	_ "github.com/oytuntez/go-words/counters/base"
	_ "github.com/oytuntez/go-words/counters/csv"
	_ "github.com/oytuntez/go-words/counters/txt"
)

func init() {}