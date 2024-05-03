package match

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Matcher struct {
	Input  io.Reader
	Output io.Writer
}

func NewMatcher() *Matcher {
	return &Matcher{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

func (m *Matcher) Match() {
	scanner := bufio.NewScanner(m.Input)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "hello") {
			fmt.Fprintf(m.Output, "%d: %s\n", lineNumber, line)
		}
		lineNumber++
	}
}
