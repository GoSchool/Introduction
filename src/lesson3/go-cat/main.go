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
	if _, err := io.Copy(dst, src); err != nil {
		fail(err)
	}
}

func determineSource(args []string) io.Reader {
	if len(args) > 0 {
		file, err := os.Open(args[0])

		if err != nil {
			fail(err)
		}

		return file
	}

	return os.Stdin
}

func fail(err error) {
	println(err.Error())
	os.Exit(1)
}
