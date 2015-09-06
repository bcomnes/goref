package main

import "fmt"

x := [5]int{1,2,3,4,5}
var total float64 = 0

for _, value := range x {
  total += value
}

fmt.Println(total / float64(len(x)))
