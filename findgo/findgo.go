package findgo

import (
	"io/fs"
	"path/filepath"
	"time"
)

func Files(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}

func OlderFiles(fsys fs.FS, duration time.Duration) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		info, err := fs.Stat(fsys, p)
		if err != nil {
			return err
		}
		if time.Since(info.ModTime()) > duration {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
