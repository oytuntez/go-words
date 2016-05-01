package registry

import (
	"testing"
	"reflect"
	"github.com/oytuntez/go-words/counters/interface"
)

type TxtTest struct {
	interfaces.CounterInterface
}

func Test(t *testing.T) {
	Add("txt", TxtTest{})
	counter := Get("txt")

	if !reflect.DeepEqual(counter, TxtTest{}) {
		t.Fail()
	}
}
