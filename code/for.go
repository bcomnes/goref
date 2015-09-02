package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }

  second()
  three()
}

func second()  {
  // A more common an concise declaration
  for i := 11; i <= 20; i++ {
    fmt.Println(i)
  }
}

func three()  {
  // A more common an concise declaration
  for i := 19; i >= 0; i-- {
    fmt.Println(i)
  }
}
