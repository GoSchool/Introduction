package main

import (
	"io"
	"os"
)

func main() {
	src := determineInputs()
	dst := os.Stdout

	cat(src, dst)
}

func cat(inputs []io.Reader, output io.Writer) {
	for _, input := range inputs {
		if _, err := io.Copy(output, input); err != nil {
			fail(err)
		}
	}
}

func determineInputs() (inputs []io.Reader) {
	inputs = append(inputs, os.Stdin)

	return
}

func fail(err error) {
	println(err)
	os.Exit(1)
}
