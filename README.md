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
- [Go by Example](https://gobyexample.com)
- [talks.golang.org](https://talks.golang.org)

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

You can quickly create new projects by creating the repo on github, then `go get [repo-qualified-name]` the newly created repo:

```sh
go get github.com/bcomnes/project-name
```

> - Go programmers typically keep all their Go code in a single workspace.
- A workspace contains many version control repositories (managed by Git, for example).
- Each repository contains one or more packages.
- Each package consists of one or more Go source files in a single directory.
- The path to a package's directory determines its import path.

Every folder can house a single package.  Sub-packages can live in subfolders of a package.

```sh
/package/
  foo.go
  bar.go
  /subpacakge/
    beep.go
    boop.go
```

Packages are a single file or collection of files that start with the `package` header.

```go
package foo
```

Every package file must have the same package name as their siblings in the folder that they live in.  Having a different package declaration name than siblings results in an error.

Every source file in your project must have a package header.

Packages that generate executable commands must always be named `package main`.

```go
package main
```

`package main` must contain a `main()` function, which is called when you execute the program.

```go
package main

import "fmt"

func notmain() {
  fmt.Println("I don't run unless called!")
}

func main() {
  // Hi I'm the entry point
  fmt.Println("Hello world")
}
```

File names are mostly irrelevant.  Packages are referenced by their directory path in `$GOPATH`:

```go
import (
  "fmt"
  "github.com/bcomnes/package-folder"
)
```

where `github.com/bcomnes/package-folder` has the following files:

```sh
/package-folder/
  package-file1.go
  package-file2.go
```

When you import a package by its dir path, the package name from the package declaration at the top of the package files becomes the prefix at which you can access everything that is exported from that package. e.g.

If `package-file1.go` from  `github.com/bcomnes/package-folder` had the following contents:

```go
//package-file1.go
package foo

import "fmt"

func Say( ){
  fmt.Println("Hi from foo")
}

```

then another go package importing this would import it like this:

```go
package bar

import "github.com/bcomnes/package-folder"

foo.Say() //Hi from foo
```

The package name `foo` is silently and invisibly dropped into the scope of the importing package as a package prefix.  You can name your imports by giving them a name upon importing

```go
package bar

import pf "github.com/bcomnes/package-folder"

pf.Say() //Hi from foo
```

Its generally a good idea to keep your package folder name the same as the package name, but know that you cannot depend on this to be true necessarily.  Many packages will leave their main package in the root of their project repo, because their primary product is a binary program installed with the `go get` command.

Packages export things to importers by capitalizing the first letter of the variable or function they are exporting.  Package files have intrinsic access to the variables and types declared anywhere else in the package siblings.  This is unfortunate, so please take this in mind and make declarations obvious, and possibly even be reluctant when creating multi-file packages.

#### More info

- [SO: What is a sensible way to layout a Go project](http://stackoverflow.com/questions/14867452/what-is-a-sensible-way-to-layout-a-go-project)

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

You can think of this as stating:

> Create a variable who's type and value is the following -- [Rob Pike: Advanced Topics in Programming Languages: Concurrency/message passing Newsqueak](https://youtu.be/hB05UFqOtFA?t=9m42s)

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

Similar to type inference, you can create an array and infer its length upon creation.  This is generally preferred over specifying length manually the same way `:=` is preferred when possible.

```go
arr := []float64{1,2,3,4}
// creates a slice with an underlying array of length 4
```

When you create an array this way, you are actually creating a slice.

Slices are 'slices' of arrays.  Slices can be shorter than their underlying array and change in length (increase and decrease) but cannot exceed the length of the underlying array.

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

TODO: grab depiction of slice smaller than underlying array.

### Slicing

Using the slice syntax, you can slice up existing arrays and create a slice of a range.

```go
// array[low:high]
arr := []float64{1,2,3,4,5,6,7,8,9}
s := arr[0:5] // [1 2 3 4 5]
t := arr[1:6] // [2 3 4 5 6]
```

The `low` index is where the slice starts and the `high` index is where the slice stops.  The `high` index is **not** included.


### Appending

The `append` function accepts a slice and additional values to append.  It returns a new slice with with those values appended.

```go
slice := []float64{1,2,3,4,5} // [1 2 3 4 5]
biggerSlice := append(slice, 6,7,8,9) // [1 2 3 4 5 6 7 8 9]
```

### Copying

The `copy` function copies as much as one slice into another slice:

```go
func copyTest() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 2)
  var slice3 []int
	copy(slice2, slice1)
  copy(slice3, slice1)
  slice2[0] = 2
	fmt.Println(slice3, slice2, slice1) // [] [2 2] [1 2 3 4]

	slice4 := slice1 // Assignments are copy by reference
	slice4[0] = 10
	fmt.Println(slice4, slice1) // [10 2 3 4] [10 2 3 4]
}
```

## Maps

AKA: "Hash Table", "Dictionary" or "Associative Array".  An unordered collection of keys-value pairs.

```go
x := make(map[string]int)
/*    ^    ^     ^    ^
      |    |     |    |
    init   |     |    |
         keyword |    |
                 |    |
                 |    |
             key type |
                      |
                 value type
*/
```

Maps need to be initialized before they can be used using the `make` keyword.  The shorthand method is preferred in general.

Functions like `len` work on maps.

Keys can be deleted using the `delete(map, keyName)` function.

Values are accessed similar to how we access arrays using the `map["key"]` syntax.

```go
package main

import "fmt"

func main() {
	x := make(map[string]string)
	x["key"] = "10"
	fmt.Println(x["key"]) // "10"
	fmt.Println(len(x)) // 1
	x["foo"] = "bar"
}
```

Accessing keys that don't exist returns the zero value of the value type:

```go
y := make(map[string]string)
z := make(map[int]int)
fmt.Println(y["foo"] == "") // true
fmt.Println(z[10] == 0) // true
```

Testing for key existence is easy due to a secondary Boolean existence return value:

```go
h := make(map[string]string)
h["beep"] = "boop"

noise, ok := h["beep"]
fmt.Println(noise, ok) // boop true
noNoise, ok  := h["foo"]
fmt.Println(noNoise, ok) // '' false
```

When used in an `if` condition:

```go
h := make(map[string]string)
h["beep"] = "boop"

if noise, ok := h["beep"]; ok {
  fmt.Println(noise) // boop
}

if noise, ok := h["bar"]; ok {
  // Doesn't run
  fmt.Println(noise)
}
```

### Map shorthand

The preferred way of making maps is with the shorthand syntax:

```go
myMap := map[string]string {
  "hey": "hi",
  "beep": "boop",
  "foo": "bar",
  "bleep": "blop", // trailing comma required
}
fmt.Println(myMap) // map[hey:hi beep:boop foo:bar bleep:blop]
```

### Maps of maps

We can create maps of maps like this:

```go
mapOfMap := map[string]map[string]int {
  "hey": {
    "beep": 10,
    "boop": 20,
  },
  "hi": {
    "beep": 30,
    "boop": 40
  }
}
fmt.Println(mapOfMap) // map[hey:map[beep:10 boop:20] hi:map[boop:40 beep:30]]
```

*Note:* Maps of maps used to require a more verbose syntax:

```go
oldMapOfMap := map[string]map[string]int {
  "hey": map[string]int{
    "beep": 10,
    "boop": 20,
  }
}
fmt.Println(mapOfMap) // map[hey:map[beep:10 boop:20] hi:map[boop:40 beep:30]]
```

This is no longer so as of (TODO: get the go version where this changed).

## Functions

A simple function example looks like:

```go
func average(xs []float64) float64 { panic("Not Implemented") }
/*^    ^     ^      ^        ^     ^
  |    |     |      |        |     |
keyword|     |      |        |     |
       |     |      |        |     |
     name    |      |        |     |
           arg1  arg-type    |     |
                             |     |
                       return-type |
                                   |
                            function-
*/
```

The combination of the function's arguments(AKA parameters) and return type is known as the **function signature**.

Calling a function in go pushes the function onto the execution callstack.  When the function returns, it is popped off the callstack and returns a value or set of values to the previous function on the callstack.

TODO: grab callstack image.

### Named Return Types

Return types can be optionally named.  Named return types are intrinsically returned when the function ends:

```go
func f2() (r int) {
  r = 1
  return
  // r is returned with whatever value is assigned to it.
}
```

### Multiple return values

Returning multiple return values is as easy as declaring them in the return section of the function declaration.

```go
func f3() (int, int) {
  return 5, 6
}

func main() {
  x, y := f3()
  fmt.Println(x, y) // 5, 6
}
```

Multiple return values are often used to indicate errors or success values:

```go
x, err := f()
x, ok := g()
```

### Variadic Functions

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
