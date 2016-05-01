package base

import "testing"

var sampleText = make(map[string]int)
var samplePaths = make(map[string]int)

func prepare() {
	sampleText["Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit..."] = 14
	samplePaths["/Users/oytun/Documents/websummit-attendees.csv"] = 92874
}

func TestCount(t *testing.T) {
	prepare()

	counter := Counter{}

	for text, wordCount := range sampleText {
		if wordCount != counter.Count(text) {
			t.Fail()
		}
	}
}

func TestText(t *testing.T) {
	prepare()

	counter := Counter{}

	for text, wordCount := range sampleText {
		if wordCount != counter.Text(text) {
			t.Fail()
		}
	}
}

func TestFile(t *testing.T) {
	prepare()

	counter := Counter{}

	for path, wordCount := range samplePaths {
		if wordCount != counter.File(path) {
			t.Fail()
		}
	}
}