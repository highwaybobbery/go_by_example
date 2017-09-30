package main

import(
  "fmt"
  "time"
  "sync/atomic"
)

func main() {
  var ops uint64 = 0

  for i := 0; i < 50; i ++ {
    fmt.Println("goRoutine:", i)
    go func() {
      for {
        atomic.AddUint64(&ops, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  fmt.Println("sleepy")
  time.Sleep(time.Second)

  fmt.Println("loady")
  opsFinal := atomiC.LoadUint64(&ops)
  fmt.Println("ops:", opsFinal)
}
