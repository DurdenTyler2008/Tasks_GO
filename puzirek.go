package main

import (
	"fmt"
)

func main() {
	var massive [10]int = [10]int{1, 5, 9, 8, 3, 16, 11, 48, 95, 21}
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if massive[j] > massive[j+1] {
				b := massive[j]
				massive[j] = massive[j+1]
				massive[j+1] = b

			}
		}
	}
	fmt.Println(massive)

}
