package main

import (
  "fmt"
)

var Number = NewType("Number", Generic)
var Integer = NewType("Integer", Number)
var Int64 = NewType("Int64", Integer)
var Int32 = NewType("Int32", Integer)
var Float = NewType("Float", Number)

func main() {
  if !DefineFunction("print", Generic) {
    panic("Couldn't define print(Generic)")
  }
  if !DefineFunction("print", Number) {
    panic("Couldn't define print(Number)")
  }
  if !DefineFunction("print", Integer) {
    panic("Couldn't define print(Integer)")
  }

  if !DefineFunction("add", Number, Number) {
    panic("Couldn't define add(Number, Number)")
  }
  if !DefineFunction("add", Integer, Integer) {
    panic("Couldn't define add(Integer, Integer)")
  }

  gen := NewVariable("gen", Generic)
  num := NewVariable("num", Number)
  i := NewVariable("i", Integer)
  i64 := NewVariable("i64", Int64)
  i32 := NewVariable("i32", Int32)
  f := NewVariable("f", Float)

  fmt.Println()
  CallFunction("print", gen)
  CallFunction("print", num)
  CallFunction("print", i)
  CallFunction("print", i64)
  CallFunction("print", i32)
  CallFunction("print", f)

  CallFunction("add", num, num)
  CallFunction("add", num, i)
  CallFunction("add", i32, i64)
  CallFunction("add", i64, i32)
  CallFunction("add", i64, f)
}
