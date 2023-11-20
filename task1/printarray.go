package main
import "fmt"
func main(){
	var array[5] string = [5]string{"apple","mango","pear","avacado","pineapple"}
	for index, value := range array {
		fmt.Println("This is ",value," and its index in the array is",index)
	}	
}