package main

import "fmt"

func sum(nums ...int) int {
  fmt.Println("nums: ", nums)
  total := 0
  for _, num := range(nums) {
    total += num
  }
  return total
}

func main() {
  nums := []int{1, 2, 3, 4}
  total := sum(nums...)
  fmt.Println("total: ", total)
}
