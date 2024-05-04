package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	files  []io.Reader
	input  io.Reader
	output io.Writer
}

type option func(*Counter) error

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c Counter) Lines() int {
	var result int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		result++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return result
}

func (c *Counter) Words() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return words
}

func (c *Counter) Bytes() int {
	bytes := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanBytes)
	for input.Scan() {
		bytes++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return bytes
}

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		if len(args) < 1 {
			return nil
		}
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func MainLines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0
}

func MainWords() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Words())
	return 0
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	byteMode := flag.Bool("bytes", false, "Count bytes, not words")
	flag.Parse()
	c, err := NewCounter(
		WithInputFromArgs(flag.Args()),
	)
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines] [bytes] [files...]\n", os.Args[0])
		fmt.Println("Counts words (or lines/bytes) from stdin (or files).")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if *lineMode && *byteMode {
		fmt.Fprintln(os.Stderr, "Cannot count lines and bytes at once.")
		return 1
	} else if *lineMode {
		fmt.Println(c.Lines())
	} else if *byteMode {
		fmt.Println(c.Bytes())
	} else {
		fmt.Println(c.Words())
	}
	return 0
}
