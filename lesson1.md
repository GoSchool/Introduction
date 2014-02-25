## Lesson 1: Getting Started

### Install

#### Instructions for Mac

*Using Homebrew (Preferred Method):*

```bash
brew install go
```

*From the distribution (use at your own risk):*

[instructions from golang.org](http://golang.org/doc/install#download)

#### Instructions for Windows

[instructions from golang.org](http://golang.org/doc/install#download)

#### Instructions for Linux

```bash
sudo apt-get update
sudo apt-get install golang
```

### Environment - GOPATH

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

### Tools

#### go fmt

#### go build

#### go test

#### go vet

### Your first program: Hello World!

Check that Go is installed correctly by building a simple program, as follows.

Create a file named hello.go and put the following program in it:

```go
package main

func main() {
    print("hello, world\n")
}
```

Then run it with the go tool:

```bash
go run hello.go
```

hello, world

If you see the "hello, world" message then your Go installation is working.

### Your first package: double


