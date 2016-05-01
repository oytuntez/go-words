package interfaces

type CounterInterface interface {
	Count(s string) int
	Text(s string) int
	File(p string) int
}
