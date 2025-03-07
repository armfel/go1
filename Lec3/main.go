package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		name    string
		surname string
		age     int
	)

	//fmt.Scan(&name)
	//fmt.Scan(&surname)
	//fmt.Scan(&age)
	//fmt.Scan(&age, &name)

	//fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

	//Для ручного использования потока
	fmt.Fscan(os.Stdin, &name)
	fmt.Fscan(os.Stdin, &surname)
	fmt.Fscan(os.Stdin, &age)
	//fmt.Println("New age:", age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS\n", name, surname, age)
}
