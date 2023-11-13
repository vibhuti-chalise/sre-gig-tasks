package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := getNumber(100, 1)
	fmt.Println("===================== Random number generated is ===================== ", n)
	checkNumber(n)
}

func getNumber(max int, min int) int {
	return rand.Intn(max - min)
}

func checkNumber(n int) {
	if n > 50 {
		if n%2 == 0 {
			fmt.Println("\n\t\t\t\t It's closer to 100, and its even!\n")
		} else {
			fmt.Println("\n\t\t\t\t It's closer to 100\n")
		}
	} else if n < 50 {
		fmt.Println("\n\t\t\t\t It's closer to 0\n")
	} else if n == 50 {
		fmt.Println("\n\t\t\t\t It's 50 \n")
	}
}
