package main

import (
	"io"
	"os"
)

func main() {
	src := determineSources(os.Args[1:])
	dst := os.Stdout

	cat(src, dst)
}

func cat(sources []io.Reader, dst io.Writer) {
	for i := range sources {
		_, err := io.Copy(dst, sources[i])
		failOnError(err)
	}
}

func determineSources(args []string) (sources []io.Reader) {
	if len(args) == 0 {
		sources = append(sources, os.Stdin)
		return
	}

	for i := range args {
		file, err := os.Open(args[i])
		failOnError(err)
		sources = append(sources, file)
	}

	return
}

func failOnError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
