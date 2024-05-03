package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestLineCounter(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("a\nb\nc\n")
	c, err := counter.NewCounter(counter.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
