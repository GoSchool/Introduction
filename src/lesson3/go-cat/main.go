package main

import (
	"io"
	"os"
)

func main() {
	src := determineSource(os.Args[1:])
	dst := os.Stdout

	cat(src, dst)
}

func cat(src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	failOnError(err)
}

func determineSource(args []string) io.Reader {
	if len(args) > 0 {
		file, err := os.Open(args[0])
		failOnError(err)
		return file
	}

	return os.Stdin
}

func failOnError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
