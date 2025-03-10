package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)
	//fmt.Println(aBoolean + aBoolean)

	//Numerics. Integers
	//int8, int16, int32, int64, int - зависит от разрядности платформы
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	//Узнаем сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент. Прииспользовании короткого объявления - тип для целого числа - int платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент 2. Используйте явное приведение типов при необходимости если уверены, что не произойдет коллизия
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперимент 3.
	// над int и intX, то обязательно нужно использовать механизм приведения. Т.к. int != int64
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64 + int64(fourthInt))
	// + - * / %

	// Аналогичнм образом устроены uint8, uint16, uint32, uint64, uint
	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of b %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum: %.2f and Sub: %.2f\n", sum, sub)

	//Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings. Строка - это набор БАЙТ
	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", name, len(name)) // Функция len() возвращает количество элементов в наборе
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))
	//Rune - руна. Это один utf-ный символ.
	//Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "cde"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' // Для инициализации руны символьным значением - используйте ''
	var thirdRune rune = 100
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)
	fmt.Println(strings.Compare("apple", "banana")) // -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte // alias uint8
	fmt.Println("Byte:", aByte)
}
