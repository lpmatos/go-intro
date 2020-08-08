# Creating and Running Simple Go Programs

## Creating Our Hello World Package

For the most part, all Go code needs to be placed within a package, and packages are named according to where the project will live in source control. We'll see a lot of projects nested under $HOME/go/src/github.com/ with a variety of different user and organization names.

To start, we're going to create a package at the root of our src directory called hello_world:

```bash
mkdir hello_world
cd hello_world
hello_world $
```

Let's create our first Go file, main.go.

```go
package main

func main() {
}
```

Technically, this is a valid program that can be built into an executable, but it won't do anything. To create an executable, we need a main function defined within the main package within a package directory. To print to the screen when we run our program, we're going to need to use the fmt package from Go's standard library. Let's modify the file to actually print a message:

```go
package main

import "fmt"

func main() {
        fmt.Println("Hello Linux Academy!")
}
```

Now that we've imported the fmt package using the import keyword, we can access functions and types from within the package by chaining off of the package's name. The Println function takes strings and prints them out with a trailing newline character.

## Running Go Code

Now that the file is written, we're ready to run our program. We're going to look at two different ways to do this:

1. Compiling and running at the same time using go run.
2. Compiling a binary to run using go build.

Note that when we use go run, we need to provide the file we want to run. The other approach, using go build, will create a binary for our program in our current directory.
