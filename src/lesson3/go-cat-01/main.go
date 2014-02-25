package main

import (
	"io"
	"os"
)

func main() {
	cat(os.Stdin, os.Stdout)
}

func cat(src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	failOnError(err)
}

func failOnError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
