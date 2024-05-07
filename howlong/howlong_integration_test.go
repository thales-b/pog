//build integration

package howlong_test

import (
	"howlong"
	"os/exec"
	"testing"
)

func TestGetTimeOutput(t *testing.T) {
	t.Parallel()
	_, err := exec.Command("/usr/bin/time", "-p", "sleep", "1").CombinedOutput()
	if err != nil {
		t.Skipf("unable to run 'time' command: %v", err)
	}
	text, err := howlong.GetTimeOutput("sleep", "1")
	if err != nil {
		t.Fatal(err)
	}
	duration, err := howlong.ParseTimeOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("(time: %fs)", duration)
}
