package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
	Transcript     io.Writer
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Stdin:      in,
		Stdout:     out,
		Stderr:     errs,
		DryRun:     false,
		Transcript: io.Discard,
	}
}

func (s *Session) Run() {
	stdout := io.MultiWriter(s.Stdout, s.Transcript)
	stderr := io.MultiWriter(s.Stderr, s.Transcript)
	fmt.Fprintf(stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Fprintln(s.Transcript, line)
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(stdout, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(stderr, "error:", err)
		}
		fmt.Fprintf(stdout, "%s> ", output)
	}
	fmt.Fprintln(stdout, "\nBe seeing you!")
}

func CmdFromString(str string) (*exec.Cmd, error) {
	args := strings.Fields(str)
	if len(args) == 0 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func Main() int {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	transcript, err := os.Create("transcript.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer transcript.Close()
	session.Transcript = transcript
	session.Run()
	return 0
}
