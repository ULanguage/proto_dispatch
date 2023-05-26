package main

import (
  "fmt"
)

type Function struct {
  Name string
  Params []*Type
  // TODO: Returns []*Type
  // NOTE: Would also have *byte to it's code
}

var Functions = []*Function{}

func NewFunction(name string, params []*Type) *Function {
  return &Function{
    Name: name,
    Params: params, // TODO: Copy?
  }
}

func (fun *Function) Call(args... *Variable) {
  fmt.Printf("Called: %v(%v) with %v\n", fun.Name, fun.Params, args)
}

func findFunction(name string, params []*Type) *Function {
  return nil
}

func DefineFunction(name string, params... *Type) bool {
  return false
}

func CallFunction(name string, args... *Variable) {
  // TODO: Return function's return
}
