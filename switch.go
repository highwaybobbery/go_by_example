package main

import "fmt"
import "time"

func main() {
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("it's morning")
  default:
    fmt.Println("it's afternoon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("boolean")
    case int:
      fmt.Println("integer")
    default:
      fmt.Printf("unknown type %T\n", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hi")
}
