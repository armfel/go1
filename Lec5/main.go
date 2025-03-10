package main

import (
	"fmt"
	"strings"
)

func main() {
	//Классический условный оператор
	//if condition {
	//	//body
	//}

	//Условный оператор с блоком else
	//if condition {

	//} else {

	//}

	var value int
	fmt.Scanln(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	//if condition1 {

	//} else if condition2 {

	//} else if ... {

	//} else {

	//}

	var color string
	fmt.Scanln(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	//Good Инициализация в блоке условного оператора
	//Блок присваявания - только :=
	//Инициализируемая переменная видна только внутри области видимости условного оператора (в телах if, else if, else)
	// Но не за его пределами
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	//Ущербно
	/*var age int = 10
	if age > 7 {
		fmt.Println("Go to school")
	} //По факту, сюда подставляется ; компилятором
	else {
		fmt.Println("Another case")
	}
	*/

	//НЕ ИДЕОМАТИЧНО
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}
	//Странное правило номер 1: в Go стараются избегать блоков else
	//Идеоматичность
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("height <= 100")
}
