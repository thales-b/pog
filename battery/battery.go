package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

func GetStatus() (Status, error) {
	text, err := GetAcpiOutput()
	if err != nil {
		return Status{}, err
	}
	return ParseAcpiOutput(text)
}

func GetAcpiOutput() (string, error) {
	data, err := exec.Command("/usr/bin/acpi").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var acpiOutput = regexp.MustCompile("([0-9]+)%")

func ParseAcpiOutput(text string) (Status, error) {
	matches := acpiOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse acp output: %q", text)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{
		ChargePercent: charge,
	}, nil
}
