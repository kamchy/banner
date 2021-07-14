package banner

import (
	"bytes"
	"testing"
)

func TestFlag(t *testing.T) {
	ifs, _ := InputFlagSet()
	var b bytes.Buffer
	ifs.SetOutput(&b)
	ifs.Usage()
	println("--- will print ---")
	if len(b.String()) == 0 {
		t.Fail()
	}
}

func TestDefaults(t *testing.T) {
	ifs, inp := InputFlagSet()
	ifs.Parse([]string{"-text", "hello"})
	if *inp.texts[0] != "hello" {
		t.Errorf("Expected text to be hello, got %s", *inp.texts[0])
	}
}
