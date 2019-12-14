package main

import "fmt"

func main() {
	// Variable examples
	// keyword variable_name type
	var i int
	i = 42

	// Assigning a value during initialization
	var j int = 43

	// Let go figure out the type
	k := 44
	fmt.Println("Default way:")
	fmt.Println(i)
	fmt.Println("Same line way:")
	fmt.Println(j)
	fmt.Println("Shortcut way:")
	fmt.Println(k)
}
