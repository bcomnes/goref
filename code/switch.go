package main

import "fmt"

func main() {
  mySwitch(0)
  mySwitch(1)
  mySwitch(2)
  strSwitch("boop")
  strSwitch("foo")
}

func mySwitch(i int) {
  switch i {
  case 0: fmt.Println("0")
  case 1: fmt.Println("2")
  case 2: fmt.Println("2")
  }
}

func strSwitch(i string) {
  switch i {
  case "boop": fmt.Println("beep")
  case "foo": fmt.Println("bar")
  }
}
