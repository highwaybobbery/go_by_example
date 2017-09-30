package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("m: ", m)

  v1 := m["k1"]
  fmt.Println("v1: ", v1)

  fmt.Println("mlen: ", len(m))

  delete(m, "k2")
  fmt.Println("m: ", m)

  _, present :=  m["k2"]

  fmt.Println("present: ", present)

  n := map[string]int { "foo": 1, "bar": 2 }
  fmt.Println("map: ", n)
  
}
