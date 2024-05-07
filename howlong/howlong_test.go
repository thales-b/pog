package howlong_test

import (
	"howlong"
	"testing"
)

func TestParseTimeOutput(t *testing.T) {
	t.Parallel()
	want := 1.00
	input := `/usr/bin/time -p sleep 1 
		real 1.00 
		user 0.00 
		sys 0.00`
	got, err := howlong.ParseTimeOutput(input)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
