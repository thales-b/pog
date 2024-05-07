package howlong_test

import (
	"howlong"
	"testing"
)

func TestParseTimeOutput(t *testing.T) {
	t.Parallel()
	want := 1.009
	input := "sleep 1  0.00s user 0.00s system 0% cpu 1.009 total"
	got, err := howlong.ParseTimeOutput(input)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
