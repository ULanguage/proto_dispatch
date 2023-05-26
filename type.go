package main

import (
  "math"
)

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

func (t1 *Type) IsDescendantOf(t2 *Type) bool {
  if t1 == t2 {
    return true
  }

  for _, p := range t1.Parents {
    if p == t2 || p.IsDescendantOf(t2) {
      return true
    }
  }

  return false
}

func (t1 *Type) DistanceTo(t2 *Type) int {
  if t1 == t2 {
    return 0
  }

  distance := math.MaxInt
  for _, p := range t1.Parents {
    newDistance := p.DistanceTo(t2) + 1
    if newDistance < distance {
      distance = newDistance
    }
  }

  return distance
}

func (t *Type) String() string {
  return t.Name
}
