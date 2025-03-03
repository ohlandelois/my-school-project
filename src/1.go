package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("The sum of 3 and 4 is", add(3, 4))
}