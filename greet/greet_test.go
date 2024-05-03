package greet_test

import (
	"bytes"
	"fmt"
	"greet"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	bufIn := bytes.NewBufferString("you")
	bufOut := new(bytes.Buffer)
	g := greet.NewGreeter()
	g.Input = bufIn
	g.Output = bufOut
	g.Greet()
	want := "What is your name? Hello, you"
	got := bufOut.String()
	if want != got {
		fmt.Errorf("want %s, got %s", want, got)
	}
}
