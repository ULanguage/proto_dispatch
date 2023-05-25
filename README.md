# Prototype for multiple-dispatch
Inspired by [Julia](https://julialang.org/) I'm testing how to make multiple-dispatch possible.

## Objective
See [this video](https://www.youtube.com/watch?v=kc9HwsxE1OY)
* Easily define new types to which existing operations apply
* Easily define new operations which apply to existing types

## Plan
My first iteration to solve this problem is to have a graph of functions organized by their name and their parameter's/return types.
