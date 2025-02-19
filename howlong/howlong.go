package howlong

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func GetTimeOutput(args ...string) (string, error) {
	cmdArgs := append([]string{"-p"}, args...)
	cmd := exec.Command("/usr/bin/time", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v", err)
	}

	return string(output), nil
}

var timeOutput = regexp.MustCompile(`real (\d+\.\d+)`)

func ParseTimeOutput(text string) (float64, error) {
	matches := timeOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return 0, fmt.Errorf("Failed to parse time output: %q", text)
	}
	duration, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse duration: %q", matches[1])
	}
	return duration, nil
}

func Main() int {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s COMMAND [args...]\n", os.Args[0])
		return 1
	}
	out, err := GetTimeOutput(os.Args[1:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	parsed, err := ParseTimeOutput(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("(time: %fs)\n", parsed)
	return 0
}
