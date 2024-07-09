# Interfaces

### What are Interfaces?
- An interface groups types together based on their methods
- For a type to "satisfy" an interface, it must have implementations of the interface associated with its type.

example:
```go
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Println("Area", s.area())
	fmt.Println("Circumference", s.circumference())
}

func main() {
	c := circle{radius: 12}
	printShapeInfo(c)
}
```

[WATCH THIS: A Practical Example How To Use Interfaces In Golang](https://www.youtube.com/watch?v=McRq-uBAa9I)