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

	//Case c множеством вариантов
	var ageGroup string = "f" //Возрастные группы: "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are too young/old")
	}

	//Case с выражениями
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too young")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you?")
	}

	//Case с проваливаниями. Проваливания выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ
	//В момент выполнения fallthrough у следующего кейса не проверяется условие
	//а сразу выполняется тело
	var number int
	fmt.Scan(&number)
outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less than 100\n", number)
		if number%2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d is GREATER than 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d is GREATER than 1000\n", number)
		fallthrough
	default:
		fmt.Printf("%d DEFAULT\n", number)
	}

	//Гадость с терминацией цикла for из switch
	var i int
underloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break underloop
		}
	}
	fmt.Println("END")
}
