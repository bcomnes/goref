# goref
A quick reference for the go language.

This started as an incomplete blogpost that never saw the light of day, so rather than let that rot, this will be a semi-living document that will be updated as needed.

Other than specific citations, most of the following is simmered down from the following sources:

- [An Introduction to Programming in Go](http://www.golang-book.com) by [Caleb Doxsey](http://www.doxsey.net)

  <a href="http://www.golang-book.com"><img src="/img/introtogo.png" width="100"></a>
- [How to Write Go Code][how-url]
- [Effective Go][effective-url]
- [Golang FAQ][faq-url]
- [Rob Pike - From Parallel to Concurrent](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2014/From-Parallel-to-Concurrent)
- [Rob Pike - Concurrency Is Not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso)

[how-url]: https://golang.org/doc/code.html
[effective-url]: https://golang.org/doc/effective_go.html
[faq-url]: https://golang.org/doc/faq

## Install

Install using a rolling package manager:

```sh
$ brew install go #OSX
```

```sh
C:> choco install golang #Windows
```

```sh
$ sudo pacman -S go #Arch Linux
```

The prebuilt binaries are also a good option for linux.  Don't get stuck using an old version.

[golang.org/dl](https://golang.org/dl/)

You also need to install `git` and `hg`.  Quite a few official packages are versioned with `hg`.

## Configure

Go expects two primary `ENV` vars to be set:

### $GOROOT

`GOROOT` needs to be set to where you installed go.  In most cases, this will automatically be set by the package manager, or if you installed the prebuilt bins to the default location.  Otherwise you need to set this in your dotfiles.  See [golang.org/doc/install#tarball_non_standard](https://golang.org/doc/install#tarball_non_standard)

### $GOPATH

The `GOPATH` path needs to be set every time unfortunately.  This seems to be the convention:

```sh
$ mkdir -p ~/go/{bin,src} ; echo "export GOPATH=\$HOME/go" >> ~/.bashrc ; echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bashrc 
```

This should add the following lines to your `.bashrc`

```sh
# add to ~/.bashrc to make the above action stik
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

The `GOPATH` enables the `go get` command which downloads and build packages from git repositories.  They get built and installed to your `GOPATH`.  You pretty much always want to run the `bin`s they come along with so adding `$GOPATH/bin` to the `PATH` is critical.

For this document assume:

```sh
$GOPATH = ~/go
```

## Basic Go Packaging

Read the entire "[How to Write Go][how-url]" document,  but here are the basics:

Develop your code in the `src` folder corresponding to where you host your code:

```sh
~/go/src/github.com/bcomnes/project-name
```

Every file in your project is a `package`.  Each package needs a package declaration at the top of the file:

```go
package foo
```

Executable commands must always run `package main`.

```go
package main
```

`package main` must contain a `main()` function, which is called when you execute the program.

```go
package main

import "fmt"

func notmain() {
  fmt.Println("I don't run!")
}

func main() {
  // Hi I'm the entry point
  fmt.Println("Hello world")
}
```

## Basic Go CLI

Run `go` files with `go run`:

```sh
$ go run /path/to/foo.go
```

Get docs using `godoc`:

```
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
```


Download and build dependencies with `go get`:

```sh
$ go get github.com/ipfs/go-ipfs/cmd/ipfs
```

`go get` has a few flags:

- `-h`: display extended help info
- `-d`: download package only (no install)
- `-t`: download package with tools necessary for testing
- `-u`: update package and dependencies

## Common Types

Go is a typed language.  You can specify to the compiler the datatype of the variable.  If you try to assign type a to a variable of type b, the compiler throws an error.

### Numbers

- `uint8` (`u` means unsigned e.g. no `+/-`)
- `uint16`
- `uint32`
- `uint64`
- `int8`
- `int16`
- `int32`
- `int64`
- `float32`
- `float64`
- `complex64`
- `complex128`

### Stings

Double quotes (`"string"`) requires escaped whitespace:

```go
x := "String \n with \n newlines and \t tabs"
```

Backticks (``string``) can contain whitespace

```go
x := `String
with
whitespaces`
```

Simple concatination of strings can be done with the `+` operator:

```go
x := "string1 "
y := "string2"
z := x + y // "string1 string2"
```

### Booleans

Same as JS:

- `&&` and
- `||` or
- `!` not

### Type Conversions

Types can be converted different types by running the variable through the desired type as a function.

For example, turning an `int` into a `float64`

```go
var x int = 32
var xFload64 = float64(x)
```

## Variables

Generally try to create varibles by inferring their type using the `:=` operator:

```go
x := "This results in x being a string"
```

Here is an example `go` program:

```go
package main

import "fmt"

func main() {
  var x string
  x = "Hi I'm a string"
  var y = "I'm an inferred string"
  z := "Beep another inferred string"
  i := "var not required!"
  fmt.Println(x) // Hi I'm a string
  fmt.Println(y) // I'm another string
  fmt.Println(z) // Beep another string
  fmt.Println(i) // var not required!
}
```

Variables are created using the `var` keyword, followed by the variable name (`x`) followed by the variable type.

`var` and the `:=` assignment operator can infer the correct type in most cases when a type is omitted.

```go
// Simple typed variables
var x string = "hello i'm variable of type string"
var y int16 = 32
var z = "var can infer type"
i := 32
```

### Naming variables

Variable names must start with a letter and can contain letters, numbers and the underscore symbole (`_`).

Unused variables throw errors and warnings.  If a variable is named (`_`), the go compiler will not complain if it is unsed.  There should never be unused named variables.

### Simple operators

Simple operators work on most variables of the same type.

- `+`: addition
- `-`: subtraction
- `*`: multiplication
- `/`: division
- `%`: modulo (remainder)
- `x += x = x + x` increment
- `x -= x = x - x` decrement
- `==` equal
- `!=` not equal

### Scope

Go is "[lexically scoped using blocks]()"[citation needed].

Variables exist inside the braces(`{}`) where they are defined, as well as in any child braces(`{}`).

A block is the code inside a pair of braces(`{}`).

### Constants

Declair constants using the `const` keyword in the same maner as the `var` keyword.

```go
const x string = "go ahead, just try to reassign me"
```

Try to avoid using constants where configuration and configurable defaults would also work.  E.G. you never need to configure something like `Pi`(π) (defined as `const` in the `math` package).

### `var` definition shorthand

The following is a common code pattern when defining variables.

```go
var (
  a = 5
  b int64 = 6
  c = "a string"
)
```

This pattern works with other keywords that are used similarly to the `var` keyword.

## Basic `fmt` usage

### Basic line printing

- `fmt.Println('print something to stdout')`: log to stdout with a newline at the end.
- `fmt.Println('hi my name is', name)`: appending string together with a space.
- `fmt.Print('no newline')`: print without a trailing newline.


### Process input

- `fmt.Scanf("%f", &var_name)`: Wait for `stdin` and write to `&var_name`.  
  - See pointers for info on the `&` prefix.

```go
package main

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
```
## Control Structures

Go has 3 simple control stuctures that can be used in many ways.

### `for` loops

Go only has 1 type of loop: the `for` loop.   As with most languages, `for` loops repeat a block of code multiple times.

```go
package main

import "fmt"

func main() {
  first()
  second()
  three()
}

func first() {
  // minimal for loop
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }
}

func second()  {
  // A more common an concise declaration
  for i := 11; i <= 20; i++ {
    fmt.Println(i)
  }
}

func three()  {
  for i := 19; i >= 0; i-- {
    fmt.Println(i)
  }
}
```

If we have an iterable variable like an array or slice we can use the `range` keyword as a shorthand in our `for` loop.

```go
package main

import "fmt"

func main() {
  x := [5]float64{1,2,3,4,5}
  var total float64 = 0

  for _, value := range x {
    total += value
  }

  fmt.Println(total / float64(len(x))) // 3
}
```

### `if` statements

If blocks work similarly to js, except there are some minor syntax differences:

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "divisible by 2")
    } else if i % 3 {
      fmt.Println(i, "divisible by 3")
    } else {
      fmt.Println(i, "beepin beep")
    }
  }
}
```

### `switch` blocks

Switches look for a matching case top down and break once a match is found.  There is no fallthrough [[Citation Needed]]().

Switches also support a default case if no match is found.

Generally avoid switch statements.

```go
package main

import "fmt"

func main() {
  mySwitch(0)
  mySwitch(1)
  mySwitch(2)
  strSwitch("boop")
  strSwitch("foo")
}

func mySwitch(i int) {
  switch i {
  case 0: fmt.Println("0")
  case 1: fmt.Println("2")
  case 2: fmt.Println("2")
  }
}

func strSwitch(i string) {
  switch i {
  case "boop": fmt.Println("beep")
  case "foo": fmt.Println("bar")
  }
}
```

## Arrays

Arrays are an ordered (numbered) sequence of elements of a single type with a fixed length.  Once created, they cannot be resized.

```go
package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100 // set 5th element (i = 4) to 100
  fmt.Println(x) // [0 0 0 0 100]
  var length = len(x)
  fmt.Println(length)
}
```

The length of an array is accessed using the built in `len(array)` function returning the length as an `int`.

Arrays can also be created with the following shorthand:

```go
x := [5]int{ 1, 2, 3, 4, 5 }
y := [5]float32 {
  98,
  93,
  80,
  52,
  56, // Trailing comma REQUIRED
  //45, // to allow commenting out elements
}
```

Use `range` in a for loop to iterate over an array:

```go
x := float64[5]
// ... assign some stuff
for i, el := range x {
  // i is index
  // el is i'th element
}
```

## Slices

Slices are 'slices' of arrays.  Slices can be shorter than their underlying array and change in length (increase and decrease) but cannot exceed the length of the underlying array.

Slices are generally preferred to using arrays in the same way `:=` is preferred when type can be inferred.  They are created using the same syntax as an array, except by omitting the length:

```go
var x []float64
// Creates a slice x of length 0
fmt.Println(len(x)) // 0
```

You can also make slices using the `make` keyword:

```go
x := make([]float64, 5)
// creates a slice with underlying array length of 5
y := make([]float64, 5, 10)
// A slice of length 5 with an underlying array of length 10
```



## Maps

## Functions

## Defers

## Panic and Recover

## Pointers

## Structs

## Interfaces

## Go routines

### Channels

## Testing

## Error handing

## Basic webservices

## Vendoring

- [<schmichael>](http://schmichael.com/) recomends https://github.com/robfig/glock for vendoring
- There is a [vendoring experiment going on in 1.5](https://golang.org/doc/go1.5) with the following [design document](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit?usp=sharing):
