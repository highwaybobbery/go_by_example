package main

import "fmt"

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  f("direct")


  go f("indirect")
  go func(msg string) {
    fmt.Println(msg)
  }("indirector")
  go f("tacos")


  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
