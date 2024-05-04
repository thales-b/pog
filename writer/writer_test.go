package writer_test

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"writer"
)

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()
	path := "testdata/write_test.txt"
	_, err := os.Stat(path)
	if err == nil {
		t.Fatalf("test artifact not cleaned up: %q", path)
	}
	defer os.Remove(path)
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
