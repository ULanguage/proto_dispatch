package main

import (
  "fmt"
)

func main() {
  if !DefineFunction("print", Generic) {
    panic("Couldn't define print for generic")
  }

  gen := NewVariable("gen", Generic)

  fmt.Println()
  CallFunction("print", gen)
}
