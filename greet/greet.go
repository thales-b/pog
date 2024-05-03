package greet

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Greeter struct {
	Input  io.Reader
	Output io.Writer
}

func NewGreeter() *Greeter {
	return &Greeter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

func (p *Greeter) Greet() {
	scanner := bufio.NewScanner(p.Input)
	fmt.Print("What is your name? ")
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(p.Output, "Hello, %s", scanner.Text())
}

func Main() {
	NewGreeter().Greet()
}
