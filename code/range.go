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
