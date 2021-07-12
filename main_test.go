package main

import (
	"testing"
)

func TestMakeStr(t *testing.T) {
	data := []struct {
		given    string
		expected string
	}{
		{given: "", expected: "ab"},
		{given: "me", expected: "ameb"},
	}

	for i, d := range data {
		if d.expected != MakeStr(d.given, SConf{"a", "b"}) {
			t.Errorf("Haha error at %d with struct %v", i, d)
		}
	}

}
