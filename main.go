package main

import (
  "fmt"
)

type Type struct {
  Name string
  Parents []*Type
}

var Generic = &Type{
  Name: "Generic",
}

func NewType(name string, parents []*Type) *Type {
  return &Type{
    Name: name,
    Parents: parents, // TODO: Copy
  }
}

type Variable struct {
  Name string
  Type *Type 
  Value int // NOTE: Would be *byte in real life
}

func NewVariable(name string, t *Type, value int) *Variable {
  return &Variable{
    Name: name,
    Type: t,
    Value: value,
  }
}

type Function struct {
  Name string
  Params []*Type
  // TODO: Returns []string
  // NOTE: Would also have *byte to it's code
}

type Node struct {
  Type string
  Funs []Function
   *Node
  Children []*Node
}

func NewFunction(name string, params []*Type) *Function {
  return &Function{
    Name: name,
    Params: params, // TODO: Copy?
  }
}

func DefineFunction(name string, params []*Type) bool {
  return false
}

func CallFunction(name string, args []*Variable) {
  // TODO: Return function's return
}

func main() {
  fmt.Println("Ahoy there!")

  if !DefineFunction("print", []*Type{Generic}) {
    panic("Couldn't define print for generic")
  }

  CallFunction("print", []*Variable{NewVariable("test", Generic, 0)})
}
