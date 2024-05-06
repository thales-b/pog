package findgo_test

import (
	"archive/zip"
	"findgo"
	"io/fs"
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestFilesCorrectlyListsFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListsFilesInZIPArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"tree/file.go",
		"tree/subfolder/subfolder.go",
		"tree/subfolder2/another.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestOlderFiles(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file1.txt": {Mode: fs.ModeAppend, ModTime: time.Now().Add(-30 * 24 * time.Hour)}, "file2.txt": {Mode: fs.ModeAppend, ModTime: time.Now().Add(-10 * 24 * time.Hour)},
		"subfolder/file3.txt":  {Mode: fs.ModeAppend, ModTime: time.Now().Add(-50 * 24 * time.Hour)},
		"subfolder/file4.txt":  {Mode: fs.ModeAppend, ModTime: time.Now().Add(-5 * 24 * time.Hour)},
		"subfolder2/file5.txt": {Mode: fs.ModeAppend, ModTime: time.Now().Add(-35 * 24 * time.Hour)},
	}
	want := []string{
		"file1.txt",
		"subfolder/file3.txt",
		"subfolder2/file5.txt",
	}
	duration := 30 * 24 * time.Hour
	got := findgo.OlderFiles(fsys, duration)
	if !cmp.Equal(want, got) {
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}
