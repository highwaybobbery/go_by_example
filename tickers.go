package main

import (
  "fmt"
  "time"
)

func main() {
  timer1 := time.NewTimer(time.Second * 2)
  done := make(chan bool)

  go func() {
    <-timer1.C
    fmt.Println("timer 1 expired")
    done <- true
  }()

  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("timer 2 expired")
  }()

  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("timer 2 stopped")
  }

  <- done
}
