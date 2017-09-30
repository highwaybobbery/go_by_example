package main

import "fmt"

type rect  struct {
  width, height int
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r rect) perim() int {
  return (r.width + r.height) * 2
}

func main() {
  box := rect{10,12}
  fmt.Println("area: ", box.area())
  fmt.Println("perim: ", box.perim())

  boxp := &box

  fmt.Println("area: ", boxp.area())
  fmt.Println("perim: ", boxp.perim())
  
}
