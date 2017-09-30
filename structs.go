package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {
  fmt.Println(person{"Bob", 23})
  fmt.Println(person{name: "Alice", age: 71})
  fmt.Println(person{name: "Fred"})
  fmt.Println(&person{name: "Ann", age: 17})

  s := person{"Tom", 50}
  fmt.Println(s.name)
  sp := &s
  fmt.Println(sp.age)
  sp.age = 39
  fmt.Println(sp)
  s.age = 12
  fmt.Println(sp)
  fmt.Println(s)
}
