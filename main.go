package main

func main() {
  if !DefineFunction("print", Generic) {
    panic("Couldn't define print for generic")
  }

  gen := NewVariable("gen", Generic)
  CallFunction("print", gen)
}
