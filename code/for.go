package main

import "fmt"

func main() {

  first()
  second()
  three()
}

func first() {
  // minimal for loop
  i := 0
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
