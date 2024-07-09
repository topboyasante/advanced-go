package main

import "fmt"

func updateName(n *string) {
	*n = " Wedge"
	fmt.Println("Name in updateName",n)
}

func main() {
	name := "Adolf"
	updateName(&name)
	fmt.Println(name)
}
