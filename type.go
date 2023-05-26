package main

type Type struct {
  Name string
  Parents []*Type
}

var Generic = NewType("Generic")

func NewType(name string, parents... *Type) *Type {
  return &Type{
    Name: name,
    Parents: parents,
  }
}
