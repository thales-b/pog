package battery_test

import (
	"battery"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseAcpiOutput_GetsChargePercent(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/acpi.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 55,
	}
	got, err := battery.ParseAcpiOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
