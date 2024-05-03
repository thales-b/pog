package match_test

import (
	"bytes"
	"fmt"
	"match"
	"testing"
)

func TestMatch(t *testing.T) {
	t.Parallel()
	bufIn := bytes.NewBufferString("hello 1\nfalse\nhello 3")
	bufOut := new(bytes.Buffer)
	m := match.NewMatcher()
	m.Input = bufIn
	m.Output = bufOut
	want := "1: hello 1\n3: hello 3"
	m.Match()
	got := bufOut.String()
	if want != got {
		fmt.Errorf("want %s, got %s", want, got)
	}
}
