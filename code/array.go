package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100 // set 5th element (i = 4) to 100
  fmt.Println(x) // [0 0 0 0 100]
  var length = len(x)
  fmt.Println(length)
}
