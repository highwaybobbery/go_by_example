package main

import "fmt"

func main() {
  messages := make(chan string, 2)

  messages <- "beef"
  messages <- "tacos"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
