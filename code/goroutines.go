package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0 ; ; i++ {
    c <- "ping" // write to channel
  }
}

func ponger(c chan string) {
  for i := 0 ; ; i++ {
    c <- "pong" // write to channel
  }
}

func printer(c chan string) {
  for {
    msg := <- c // block this go routine waiting for messages
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string // blocks main from ending till a character is typed
  fmt.Scanln(&input)
}
