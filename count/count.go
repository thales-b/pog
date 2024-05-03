package count

import (
	"bufio"
	"errors"
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

func Main() int {
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
