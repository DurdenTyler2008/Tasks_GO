package main

import "fmt"

func add(x int, y int) int {
	return x + y}
func add(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println("rezult1=", add(42, 13))
	//fmt.Println("rezult2=", add(30.0, 15.0))
}
