package banner

import (
	"bytes"
	"strconv"
	"strings"
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

type StrOp string
type IntOp int

type InputValue interface {
	Stringer
	Check(i Input, f func(Input) InputValue, exp InputValue) bool
}

func (s StrOp) Check(i Input, f func(Input) InputValue, exp InputValue) bool {
	return f(i) == exp
}

type Stringer interface {
	String() string
}

func (s StrOp) String() string {
	return string(s)
}

func (s IntOp) Check(i Input, f func(Input) InputValue, exp InputValue) bool {
	return f(i) == exp
}
func (i IntOp) String() string {
	return strconv.Itoa(int(i))
}

func TestWithLines(t *testing.T) {
	data := []struct {
		s string
		i InpData
	}{
		{"", InpData{5, 600, "out.png", 0, "", "", 30.0, 800}},
		{"-alg|1", InpData{1, 600, "out.png", 0, "", "", 30.0, 800}},
		// FIXME needs clamping
		{"-alg|100", InpData{DEF_ALG, 600, "out.png", 0, "", "", 30.0, 800}},
		{"-st|aaa", InpData{5, 600, "out.png", 0, "", "aaa", 30.0, 800}},
		{"-t|foo bar", InpData{5, 600, "out.png", 0, "foo bar", "", 30.0, 800}},
		{"-t|foo bar|-st|hello world", InpData{5, 600, "out.png", 0, "foo bar", "hello world", 30.0, 800}},
	}

	for _, d := range data {
		ifs, inp := InputFlagSet()
		ifs.Parse(strings.Split(d.s, "|"))
		inp.Clamp()
		got := new(InpData).From(inp)
		if !(got == d.i) {
			t.Errorf("Got, expected:\n%+v\n%+v", got, d.i)

		}
	}
}

func TestDefaults(t *testing.T) {
	data := []struct {
		opt      string
		optval   InputValue
		fn       func(Input) InputValue
		expected InputValue
	}{
		{"-t", StrOp("hello"), func(i Input) InputValue { return StrOp(*i.Texts[0]) }, StrOp("hello")},
		{"-st", StrOp("world"), func(i Input) InputValue { return StrOp(*i.Texts[1]) }, StrOp("world")},
		{"-o", StrOp("out.png"), func(i Input) InputValue { return StrOp(*i.OutName) }, StrOp("out.png")},
		{"-alg", IntOp(1), func(i Input) InputValue { return IntOp(*i.AlgIdx) }, IntOp(1)},
	}

	for _, d := range data {
		ifs, inp := InputFlagSet()
		ifs.Parse([]string{d.opt, Stringer(d.optval).String()})
		inp.Clamp()
		if d.fn(inp) != d.expected {
			t.Errorf("Expected %v to be %v, got %s", d.opt, d.expected, d.fn(inp))
		}
	}

}
