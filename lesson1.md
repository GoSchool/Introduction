# Introduction to Go - Lesson 1: Getting Started

## Installing Go

### Instructions for Mac

*Using Homebrew (Preferred Method):*

```bash
brew install go
```

*From the distribution (use at your own risk):*

[instructions from golang.org](http://golang.org/doc/install#download)

### Instructions for Windows

[instructions from golang.org](http://golang.org/doc/install#download)

### Instructions for Linux

```bash
sudo apt-get update
sudo apt-get install golang
```

## Environment - GOPATH

The GOPATH environment variable specifies the location of your workspace. It is likely the only environment variable you'll need to set when developing Go code.

To get started, create a workspace directory and set GOPATH accordingly. Your workspace can be located wherever you like, but we'll use $HOME/go in this document. Note that this must not be the same path as your Go installation.

```bash
mkdir $HOME/go
export GOPATH=$HOME/go
```

For convenience, add the workspace's bin subdirectory to your PATH:

```bash
export PATH=$PATH:$GOPATH/bin
```

## Tools

Difference between the naked call:

```bash
go fmt
```

And the recursive call
```bash
go fmt ./...
```

### go fmt

Usage:

```bash
go fmt [-n] [-x] [packages]
```

Fmt runs the command 'gofmt -l -w' on the packages named by the import paths.
It prints the names of the files that are modified.

For more about gofmt, see 'godoc gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

Simple usage:

```bash
go fmt
```

### go build

For an extensive reference, see the [doc](http://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

From the root of your package:

```bash
go build
```

### go test

For an extensive reference, see the [doc](http://golang.org/cmd/go/#hdr-Test_packages)

From the root of your package:

```bash
go test
```

### go vet

Usage:

```bash
go vet [-n] [-x] [packages]
```

Vet runs the Go vet command on the packages named by the import paths.

For more about vet, see 'godoc code.google.com/p/go.tools/cmd/vet'.
For more about specifying packages, see 'go help packages'.

To run the vet tool with specific options, run 'go tool vet'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

Simple usage:

```bash
go vet
```

## Your first program: Hello World!

Check that Go is installed correctly by building a simple program, as follows.

Create a file named hello.go and put the following program in it:

```go
package main

func main() {
    println("hello, world")
}
```

Then run it with the go tool:

```bash
go run hello.go
```

hello, world

If you see the "hello, world" message then your Go installation is working.

## Your first package: double

Let's get to work!

