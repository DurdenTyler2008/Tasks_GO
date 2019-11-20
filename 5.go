package main

import "fmt"

func add(x int, y int, z int) int {
	return x + y - z
}

func main() {
	fmt.Println("rezult1=", add(42, 13, 5))
	fmt.Println("rezult2=", add(30.0, 15.0, 52.0))
}
