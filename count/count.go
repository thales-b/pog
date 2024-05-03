package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	inputs []io.Reader
	output io.Writer
}

type option func(*Counter) error

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		inputs: []io.Reader{os.Stdin},
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
	for _, input := range c.inputs {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			result++
		}
	}
	return result
}

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.inputs = append(c.inputs, input)
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		if len(args) < 1 {
			return nil
		}
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				return err
			}
			c.inputs = append(c.inputs, f)
		}
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
