package main

import "fmt"

func main() {
	//Declare an array of 4 elements
	var menu [4]string

	//Initialize the first elements of the array
	menu = [4]string{"Burrito", "Pasta"}

	//Add the second, third elements to an array
	menu[3] = "Salad"
	menu[2] = "Hamburger"

	//Pass array elements by reference during a function call
	displayItem(&menu)
}

// Use a pointer variable to string array to point the address of each array items
func displayItem(food *[4]string) {
	for _, value := range food {
		fmt.Println("\tFood:", value)
	}
}
