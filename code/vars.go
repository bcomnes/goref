package main

import "fmt"

var (
  a = 5
  b int64 = 6
  c = "a string"
)

func main() {
  var x string
  x = "Hi I'm a string"
  var y = "I'm another string"
  z := "Beep another string"
  i := "var not required!"
  fmt.Println(x) // Hi I'm a string
  fmt.Println(y) // I'm another string
  fmt.Println(z) // Beep another string
  fmt.Println(i) // var not required!
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
