package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestLineCounter(t *testing.T) {
	t.Parallel()
	c := counter.NewCounter()
	c.Input = bytes.NewBufferString("a\nb\nc\n")
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
