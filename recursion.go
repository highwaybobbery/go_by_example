package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n - 1)
}

func main() {
  nums := []int{1,4,9,16}
  for _, num := range(nums) {
    fmt.Printf("n: %d fact: %d\n", num, fact(num))
  }
}
