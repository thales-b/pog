package main

import (
	"findgo"
	"fmt"
	"os"
)

func main() {
	paths := findgo.Files(os.DirFS(os.Args[1]))
	for _, p := range paths {
		fmt.Println(p)
	}
}
