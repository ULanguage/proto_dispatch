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

func filterByName(options []*Function, name string) []*Function {
  //fmt.Println("[filterByName]", name)
  newOptions := make([]*Function, 0, len(options))

  for _, fun := range options {
    if fun.Name == name {
      newOptions = append(newOptions, fun)
    }
  }

  return newOptions
}

func filterByLen(options []*Function, l int) []*Function {
  //fmt.Println("[filterByLen]", l)
  newOptions := make([]*Function, 0, len(options))

  for _, fun := range options {
    if len(fun.Params) == l {
      newOptions = append(newOptions, fun)
    }
  }

  return newOptions
}

func filterByArg(options []*Function, p *Type, idx int) []*Function {
  //fmt.Println("[filterByArg]", p, idx)
  newOptions := make([]*Function, 0, len(options))

  for _, fun := range options {
    if p.isDescendantOf(fun.Params[idx]) {
      newOptions = append(newOptions, fun)
    }
  }

  return newOptions
}

func deepestOption(options []*Function) *Function {
  if len(options) == 0 {
    return nil
  }

  return options[0] // TODO: Actually deepest
}

func findFunction(name string, params []*Type) *Function {
  options := Functions
  //fmt.Println(options)

  options = filterByName(options, name)
  //fmt.Println(options)

  options = filterByLen(options, len(params))
  //fmt.Println(options)

  for i, p := range params {
    options = filterByArg(options, p, i)
    //fmt.Println(options)
  }

  return deepestOption(options)
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
  types := make([]*Type, 0, len(args))
  for _, arg := range args {
    types = append(types, arg.Type)
  }

  fun := findFunction(name, types)
  if fun == nil {
    panic("No such function")
  }

  fun.Call(args...)
}
