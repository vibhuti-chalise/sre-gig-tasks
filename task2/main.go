package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	checkNumber(n)
}

func checkNumber(n int) {
	if n > 50 {
		if n%2 == 0 {
			fmt.Println("\t\t It's closer to 100, and its even!")
		} else {
			fmt.Println("\t\t It's closer to 100")
		}
	} else if n < 50 {
		fmt.Println("\t\t It's closer to 0")
	} else if n == 50 {
		fmt.Println("\t\t It's 50")
	}
	fmt.Println("=========== Random number generated is ============== ", n)
}
