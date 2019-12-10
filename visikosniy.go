package main

import (
	"fmt"
	"os"
)

func main() {
	var year uint16 // 2 байта от 0 до 65535
	fmt.Print("Введите год: ")
	fmt.Fscan(os.Stdin, &year)

	//fmt.Println(year)
	if year%4 == 0 {
		fmt.Println("Високосный")
	} else if year%100 != 100 {
		fmt.Println("Обычный")
	} else if year%400 == 0 {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Обычный")
	}

}
