package main

import "fmt"

func main() {
	var menu []string = []string{"Burrito", "Pasta"}

	displayItem(menu)
}

func displayItem(food []string) {
	food = append(food, "Hamburger")
	food = append(food, "Salad")
	for _, value := range food {
		fmt.Println("\tFood:", value)
	}
}
