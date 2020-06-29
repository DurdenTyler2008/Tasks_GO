// написать код, в котором человек с 3х попыток угадывает загаданное число от 0 до 9

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0
	var guess int = -1
	rand.Seed(time.Now().UnixNano())
	var number int = rand.Intn(10)

	for count < 3 && guess != number {
		fmt.Print("Введите номер от 0 до 9: ")
		fmt.Scanln(&guess)
		if guess != number {
			if guess < number {
				fmt.Println("Ваше число мньше загаданного")
			} else {
				fmt.Println("Ваше число больше загаданного")
			}
			count++
		}
	}
	if guess == number {
		fmt.Println("Вы молодец, угадали!")
	} else {
		fmt.Println("Вы не угадали,попробуйте еще раз))")
	}

}
