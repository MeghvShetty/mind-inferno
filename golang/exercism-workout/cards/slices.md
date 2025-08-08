___
Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

```go
var empty []int // an empty slice 
withData := []int{0,1,2,3,4,5} // a slice pre-filled with some data
```

You can get or set an element at a given zero-based index using square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements. Once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index. If you don't specify a starting index, it defaults to 0. If you don't specify an ending index, it defaults to the length of the slice.

```go 
newSlice := withData[2:4] 
// => []int{2,3} 
newSlice := withData[:2] 
// => []int{0,1} 
newSlice := withData[2:] 
// => []int{2,3,4,5} 
newSlice := withData[:] 
// => []int{0,1,2,3,4,5}
```

You can add elements to a slice using the `append` function. Below we append `4`and `2` to the `a` slice.

```go
a := []int{1, 3} 
a = append(a, 4, 2) 
// => []int{1,3,4,2}
```

### Indexes in slices

Working with indexes of slices should always be protected in some way by a check that makes sure the index actually exists. Failing to do so will crash the entire application.

### Empty slices

`nil`-slices are the default empty slice. They have no drawbacks towards a slice with no values in them. The `len` function works on `nil`-slices, items can be added without initializing it, and so on. If creating a new slice prefer `var s []int`(`nil`-slice) over `s := []int{}` (empty, non-`nil` slice).

### Performance

When creating slices to be filled iteratively, there is a low hanging fruit to improve performance, if the final size of the slice is known. The key is to minimize the amount of times memory has to be allocated, which is rather expensive and happens if the slice grows beyond its allocated memory space. The safest way to do this is to specify a capacity `cap` for the slice with `s := make([]int, 0, cap)` and then to `append` to the slice as usual. This way the space for `cap` amount of items is allocated immediately while the slice length is zero. In practice, `cap` is often the length of another slice: `s := make([]int, 0, len(otherSlice))`.

### Append is not a pure function

The append function of Go is optimised for performance and therefore does not make a copy of the input slice. This means that the original slice (1st parameter in `append`) will be changed sometimes.

[Slice go by example](https://gobyexample.com/slices)
[Go Github](https://github.com/golang/go/wiki/SliceTricks)
[Go-dev](https://go.dev/blog/slices-intro)
[Go-Dev Programming language ](https://go.dev/ref/spec#Slice_types)

# Variadic Functions
---
### About Variadic Functions

Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.
``
```
func find(a int, b ...int) { // ... }
```

```go
package main

  

import "fmt"

  

func find(num int, nums ...int) {

fmt.Printf("type of nums is %T\n", nums)

  

for i, v := range nums {

if v == num {

fmt.Println(num, "found at index", i, "in", nums)

return

}

}

fmt.Println(num, "not found in ", nums)

}

  
func main() {

find(89, 90, 91, 95)

// Output:

// type of nums is []int

// 89 not found in [90 91 95]

find(45, 56, 67, 45, 90, 109)

// Output:

// type of nums is []int

// 45 found at index 2 in [56 67 45 90 109]

find(87)

// Output:

// type of nums is []int

// 87 not found in []

}
```

Sometimes you already have a slice and want to pass that to a variadic function. This can be achieved by passing the slice followed by `...`. That will tell the compiler to use the slice as is inside the variadic function. The step described above where a slice is created will simply be omitted in this case.

```go
list := []int{1,2,3}
find(1, list...) // "find" defined as shown above
```

[Variadic ](https://gobyexample.com/variadic-functions)
[Function types](https://go.dev/ref/spec#Function_types)
[Variadic functions in Go](https://medium.com/rungo/variadic-function-in-go-5d9b23f4c01a)

==Caution==
==The variadic parameter must be the last parameter of the function.==
