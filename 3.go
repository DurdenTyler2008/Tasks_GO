//генератор случайного ответа

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	answers := []string{
		"Уфа",
		"Питер",
		"Самара",
		"Смоленск",
		"Тверь",
	}
	fmt.Println("Столица России:", answers[rand.Intn(len(answers))])
}
