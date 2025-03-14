package main

import "fmt"

func main() {
	//Switch!
	var price int
	fmt.Scan(&price)

	// В switch-case запрещены дублирующиеся условия в case'ах
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		//Отрабатывает, если ни один из вышеперечисленных кейсов - не сработал
		fmt.Println("Default case")
	}
}
