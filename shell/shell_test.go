package shell_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"shell"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	want := exec.Command("echo", "Hello", "world").Args
	cmd, err := shell.CmdFromString("echo Hello world")
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNewSession_CreatesExpectedSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		DryRun:     false,
		Transcript: io.Discard,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRunProducesExpectedOutput(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	out := new(bytes.Buffer)
	session := shell.NewSession(in, out, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRunCreatesTranscript(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	transcript := new(bytes.Buffer)
	session := shell.NewSession(in, io.Discard, io.Discard)
	session.DryRun = true
	session.Transcript = transcript
	session.Run()
	want := "> echo hello\necho hello\n> \n> \nBe seeing you!\n"
	got := transcript.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
