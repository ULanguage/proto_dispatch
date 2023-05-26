package main

import (
  "fmt"
)

var Number = NewType("Number", Generic)

func main() {
  if !DefineFunction("print", Generic) {
    panic("Couldn't define print for generic")
  }

  gen := NewVariable("gen", Generic)
  num := NewVariable("num", Number)

  fmt.Println()
  CallFunction("print", gen)
  CallFunction("print", num)
}
