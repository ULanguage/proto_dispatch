package main

import (
  "fmt"
)

type Variable struct {
  Name string
  Type *Type 
  // NOTE: Would also have *byte in real life
}

func NewVariable(name string, t *Type) *Variable {
  return &Variable{
    Name: name,
    Type: t,
  }
}

func (v *Variable) String() string {
  return fmt.Sprintf("(%v, %v)", v.Name, v.Type)
}
