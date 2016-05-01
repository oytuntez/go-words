package interfaces

// Interface for all counters. They all must implement these methods.
// Most of them will have these methods embedded by Counter struct.
type CounterInterface interface {
	Count(s string) int
	Text(s string) int
	File(p string) int
}
