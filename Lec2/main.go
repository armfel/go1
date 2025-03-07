package main

import (
	"fmt"
	"math"
)

func main() {
	//Простейший вывод на консоль. println - это вывод + '\n'
	fmt.Println("Hello, World!", "Hello, another!")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//Форматироанный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 43)
	////////////////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)
	//Декларирование и инициализация пользовательских значений
	var height int = 185
	fmt.Println("My heigh is:", height)
	//"Полустрогость" типизации. Можно опускать тип перменной
	var uid = 12345
	fmt.Println("My uuid:", uid)
	fmt.Printf("%T\n", uid)
	//Декларирование и инициализация переменных одного типа
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)
	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUID)

	//Немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b)
	a = 300
	//Немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200

	//Короткое объявление
	count := 30
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)

	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	println(aArg, bArg)
	aArg, bArg = 30, 40
	println(aArg, bArg)
	//aArg, bArg := 10, 30
	//println(aArg, bArg)

	//Исключения
	bArg, cArg := 300, 400
	fmt.Println(bArg, cArg)

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is: %.2f\n", math.Min(width, length))

	var (
		word1 string = "имя"
		word2 string = "твое"
		word3 string = "мне"
		word4 string = "знакомо"
	)
	fmt.Printf("%s %s %s %s\n%s %s %s %s\n%s %s %s %s\n", word4, word3, word2, word1, word3, word1, word4, word2, word2, word4, word1, word3)
}
