# Pointers

### What are pointers?
- Pointers are variables that hold the memory address of another variable.
- Because they are variables, they are also stored somewhere in memory and have a memory address
- Pointers are useful when you want to refer to a variable without copying its value.

- If x is a variable, where `x := &variable`, then `x` is a pointer that holds the memory address of the variable `variable`. 
- The type of `x` is now `*type`, where `type` is the data type of `variable`.
- If y is a variable, where `y := *x`, then `y` would hold the value in the memory address that `x` is pointing to


EXAMPLE WITHOUT POINTER: 
```go
package main

import "fmt"

func updateName(n string) {
	n = " Wedge"
	fmt.Println("Name in updateName",n) // Prints "Wedge"
}

func main() {
	name := "Adolf"
	updateName(name)

	fmt.Println(name) //Prints "Adolf"
}
```

WITH POINTER

```go
package main

import "fmt"

// This function is now saying "Hey, give me the memory address of the variable you want to change."
func updateName(n *string) {
	//Here, it says that "The value stored at the memory location provided will now be changed to "Wedge"
	*n = " Wedge"
	fmt.Println("Name in updateName",n)
}

func main() {
	name := "Adolf"
	updateName(&name)
	fmt.Println(name)
}

```