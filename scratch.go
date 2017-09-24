package main

import "fmt"

func main() {
  for n := 0; n <= 9; n++ {
    if n%2 == 0 {
        continue
    } else if n == 7 {
        break
    } else {
      fmt.Println("n:", n)
    }
  }
}
