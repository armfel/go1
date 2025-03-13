package main

import "fmt"

func main() {
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если верно - то тело цикла выполняется, если нет - то цикл завершается)
	// post - изменение переменной цикла (инкрементарное, декрементальное действие)
	// }

	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}

	//Важный момент: в качестве init может быть использовано выражение присваивания ТОЛЬКО через :=

	//break - команда, прерывающая текущее выполнение тела цикла и передающая управление инструкциям, следующим за циклом
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	//continue - команда, прерывающая текущее выполнение тела цикла и передающая управление СЛЕДУЮЩЕЙ ИТЕРАЦИИ ЦИКЛА
	for i := 0; i < 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with CONTINUE")

	//Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идее выше треугольник")

	//Иногда бывает плохо. С лейблами лучше
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу чтобы вообще все циклы (внешние тоже остановились)
			}
		}
	}
}
