package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() *Counter {
	return &Counter{
		Input: os.Stdin,
	}
}

func (c Counter) Count() int {
	var result int
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		result++
	}
	return result
}

func Main() {
	fmt.Println(NewCounter().Count())
}
