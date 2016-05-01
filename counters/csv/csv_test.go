package csv

import "testing"

var samplePaths = make(map[string]int)

func prepare() {
	samplePaths["/Users/oytun/Documents/websummit-attendees.csv"] = 92874
}

func TestFile(t *testing.T) {
	prepare()

	counter := CSV{}

	for path, wordCount := range samplePaths {
		if wordCount != counter.File(path) {
			t.Fail()
		}
	}
}