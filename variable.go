package main

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
