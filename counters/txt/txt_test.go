package txt

import "testing"

var samplePaths = make(map[string]int)

func prepare() {
	samplePaths["/Users/oytun/Documents/log-excerpt.log"] = 10416
}

func TestFile(t *testing.T) {
	prepare()

	counter := Txt{}

	for path, wordCount := range samplePaths {
		if wordCount != counter.File(path) {
			t.Fail()
		}
	}
}