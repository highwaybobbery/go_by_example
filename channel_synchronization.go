package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Println("working time")
  time.Sleep(time.Second)
  fmt.Println("sleepy time")

  done <- true
}

func main() {

  done := make(chan bool, 1)

  go worker(done)

  <-done

}
