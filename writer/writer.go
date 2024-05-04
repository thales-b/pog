package writer

import (
	"flag"
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}

func Main() int {
	sizePtr := flag.Int("size", 0, "size of the file in bytes")
	flag.Usage = func() {
		fmt.Printf("Usage: %s -size <num> <filename>\n", os.Args[0])
		fmt.Println("Creates a file with configurable size filled with zeros.")
	}
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		flag.Usage()
		return 1
	}
	size := *sizePtr
	content := make([]byte, size)
	err := WriteToFile(filename, content)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return 1
	}
	return 0
}
