# goref
A quick reference for to the go language.

This started as an incomplete blogpost that never saw the light of day, so rather than let that rot, this will be a semi-living document that will be updated as needed.

Other than specific citations, most of the following is simmered down from the following sources:

- [An Introduction to Programming in Go](http://www.golang-book.com) by [Caleb Doxsey](http://www.doxsey.net)

  <a href="http://www.golang-book.com"><img src="http://www.golang-book.com/assets/img/cover.png" width="100"></a>
- [How to Write Go Code][how-url]
- [Effective Go][effective-url]
- [Golang FAQ][faq-url]


## Install

Install using a rolling package manager:

{% highlight sh %}
$ brew install go #OSX
$ choco install golang #Windows
{% endhighlight %}

The prebuilt binaries are also a good option for linux.  Don't get stuck using an old version.

[golang.org/dl](https://golang.org/dl/)

You also need to install `git` and (*sigh*) `hg`.

## Configure

Go expects two primary `ENV` vars to be set:

### $GOROOT

`GOROOT` needs to be set to where you installed go.  In most cases, this will automatically be set by the package manager, or if you installed the prebuilt bins to the default location.  Otherwise you need to set this in your dotfiles.  See [#tarball_non_standard](https://golang.org/doc/install#tarball_non_standard)

### $GOPATH

The `GOPATH` path needs to be set every time unfortunately.  This seems to be the convention:

{% highlight sh %}
$ mkdir -p ~/go/{bin,src} ; export GOPATH=$HOME/go ; export PATH=$PATH:$GOPATH/bin
{% endhighlight %}

The `GOPATH` enables the `go get` command which downloads and build packages from git repositories.  They get built and installed to your `GOPATH`.  You pretty much always want to run the `bin`s they come along with so adding `$GOPATH/bin` to the `PATH` is critical.

For this document assume:

{% highlight sh %}
$GOPATH = ~/go
{% endhighlight %}

## Basic Go Packaging

Read the entire "[How to Write Go][how-url]" document,  but here are the basics:

Develop your code in the `src` folder corresponding to where you host your code:

{% highlight sh %}
~/go/src/github.com/bcomnes/project-name
{% endhighlight %}

Every file in your project is a `package`.  Each package needs a package declaration at the top of the file:

{% highlight go %}
package foo
{% endhighlight %}

Executable commands must always run `package main`.

{% highlight go %}
package main
{% endhighlight %}

`package main` must contain a `main()` function, which is called when you execute the program.

{% highlight go %}
package main

import "fmt"

func notmain() {
  fmt.Println("I don't run!")
}

func main() {
  // Hi I'm the entry point
  fmt.Println("Hello world")
}
{% endhighlight %}

## Basic Go CLI

Run `go` files with `go run`:

{% highlight sh %}
$ go run /path/to/foo.go
{% endhighlight %}

Get docs using `godoc`:

{% highlight sh %}
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
{% endhighlight %}

Download and build dependencies with `go get`:

{% highlight sh %}
$ go get github.com/ipfs/go-ipfs/cmd/ipfs
{% endhighlight %}

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

{% highlight go %}
x := "String \n with \n newlines and \t tabs"
{% endhighlight %}

Backticks (``string``) can contain whitespace

{% highlight go %}
x := `String
with
whitespaces`
{% endhighlight %}

Simple concatination of strings can be done with the `+` operator:

{% highlight go %}
x := "string1 "
y := "string2"
z := x + y // "string1 string2"
{% endhighlight %}

### Booleans

Same as JS:

- `&&` and
- `||` or
- `!` not

## Variables

Generally try to create varibles by inferring their type using the `:=` operator:

{% highlight go %}
x := "This results in x being a string"
{% endhighlight %}

Here is an example `go` program:

{% highlight go %}
package main

import "fmt"

func main() {
  var x string
  x = "Hi I'm a string"
  fmt.Println(x) // "Hi I'm a string"
}
{% endhighlight %}

Variables are created using the `var` keyword, followed by the variable name (`x`) followed by the variable type.


[how-url]: https://golang.org/doc/code.html
[effective-url]: https://golang.org/doc/effective_go.html
[faq-url]: https://golang.org/doc/faq
