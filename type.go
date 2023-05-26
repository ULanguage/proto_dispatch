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

func (t1 *Type) isDescendantOf(t2 *Type) bool {
  if t1 == t2 {
    return true
  }

  for _, p := range t1.Parents {
    if p == t2 || p.isDescendantOf(t2) {
      return true
    }
  }

  return false
}

func (t *Type) String() string {
  return t.Name
}
