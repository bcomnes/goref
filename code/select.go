package main

import (
  "fmt"
  "time"
)

func sender1(c chan<- string) {
  for {
    c <- "from 1"
    time.Sleep(time.Second * 2)
  }
}

func sender2(c chan<- string) {
  for {
    c <- "from 2"
    time.Sleep(time.Second * 3)
  }
}

func selector(chan1 <-chan string, chan2 <-chan string) {
  for {
    select {
    case msg1 := <- chan1:
      fmt.Println(msg1)
    case msg2 := <-chan2:
      fmt.Println(msg2)
    case <- time.After(time.Second):
      fmt.Println("timeout")
    }
  }
}

func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go sender1(c1)
  go sender2(c2)
  go selector(c1, c2)

  var input string
  fmt.Scanln(&input)
}
