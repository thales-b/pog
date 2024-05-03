package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
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

func (c Counter) Count() int {
	var result int
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		result++
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

func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Count())
}
