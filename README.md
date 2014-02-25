# Introduction to Go

A fast paced, hands on workshop on Go Programming. It's meant to be
carried out in 2 hours.

* Workshop Outline
	* Welcome                             | 15 minutes
	* Lesson 1: Getting Started           | 15 minutes
	* Lesson 2: Language Fundamentals     | 30 minutes
	* Lesson 3: Command Line Applications | 60 minutes


## Welcome

* Why Go?
* About us

## Lesson 1: Getting Started

* Install

> Instructions for Mac
  Using Homebrew (Preferred Method):

  ```bash
  brew install go
  ```

  From the distribution (use at your own risk):

  [instructions from golang.org](http://golang.org/doc/install#download)

> Instructions for Windows

  You can install go on windows either using an MSI installar or a zip package located here:

  [instructions from golang.org](http://golang.org/doc/install#download)

> Instructions for Linux

  ```bash
  sudo apt-get update
  sudo apt-get install golang
  ```

* Environment
	* GOPATH
* Tools
  * go fmt
	* go build
	* go test
	* go vet
* Your first program: Hello World! [example](http://play.golang.org/p/aGiaVaMIUFa)
* Your first package: double


## Lesson 2: Language Fundamentals

* Constants, Variables, Basic Types and Values

* Conditionals
	* If/Else

	* Switch
		* switch value { }
		* switch { case expression: ... }

* Collections
	* Arrays
	* Slices
	* Maps

* For
	Traditional loop for i:=0;... {
	Infinite Loop  for {
	Conditional Loop for <bool> {
	Range Loop

* Functions
	Double(int) int
	Triple(int) int
	anonymous

* Scope

* Goroutines

* Channels


## Lession 3: Command Line Applications
	* `cat`- io.Reader and io.Writer



Prequisites...
Experienced programmers, comfortable with a command line interface.
