//go:build !integration

package battery_test

import (
	"battery"
	"bytes"
	"os/exec"
	"testing"
)

func TestGetAcpiOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	data, err := exec.Command("/usr/bin/acpi").
		CombinedOutput()
	if err != nil {
		t.Skipf("unable to run 'acpi' command: %v", err)
	}
	if !bytes.Contains(data, []byte("Battery")) {
		t.Skip("no battery fitted")
	}
	text, err := battery.GetAcpiOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParseAcpiOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
