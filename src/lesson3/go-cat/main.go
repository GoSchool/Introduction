package main

import (
	"io"
	"os"
)

func main() {
	src := determineInput(os.Args)
	dst := os.Stdout

	cat(src, dst)
}

func cat(input io.Reader, output io.Writer) {
	_, err := io.Copy(output, input)

	if err != nil {
		fail(err)
	}
}

func determineInput() (input io.Reader) {


	return os.StdIn
}

func fail(err error) {
	println(err)
	os.Exit(1)
}
