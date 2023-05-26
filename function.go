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

func (fun *Function) Call(args... *Variable) {
  fmt.Printf("Called: %v(%v) with %v\n", fun.Name, fun.Params, args)
}

func findFunction(name string, params []*Type) *Function {
  return nil
}

func DefineFunction(name string, params... *Type) bool {
  if findFunction(name, params) != nil {
    return false
  }

  fun := &Function{
    Name: name,
    Params: params,
  }

  Functions = append(Functions, fun)
  return true
}

func CallFunction(name string, args... *Variable) {
  types := make([]*Type, len(args))
  for _, arg := range args {
    types = append(types, arg.Type)
  }

  fun := findFunction(name, types)
  if fun == nil {
    panic("No such function")
  }

  fun.Call(args...)
}
