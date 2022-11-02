package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)
	var num = []int{num1, num2, num3}
	for _, elm := range num {
		if elm == 0 {
			fmt.Println("Zero")
		}
		if elm == 1 {
			fmt.Println("One")
		}
		if elm == 2 {
			fmt.Println("Two")
		}
		if elm == 3 {
			fmt.Println("Three")
		}
		if elm == 4 {
			fmt.Println("Four")
		}
		if elm == 5 {
			fmt.Println("Five")
		}
		if elm == 6 {
			fmt.Println("Six")
		}
		if elm == 7 {
			fmt.Println("Seven")
		}
		if elm == 8 {
			fmt.Println("Eight")
		}
		if elm == 9 {
			fmt.Println("Nine")
		}
		if elm == 10 {
			fmt.Println("Ten")
		}
	}
}
