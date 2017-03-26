package main

import "fmt"

func makeEvenGenerator() func() uint {
  i := uint(0)

  even := func () (ret uint) {
    ret = i
    i += 2
    return
  }

  return even
}

func main() {
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}
