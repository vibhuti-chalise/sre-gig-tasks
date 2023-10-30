import "fmt"
func main(){
	var slice[] string
	slice = append(slice,"Hamburger")
	slice = append(slice, "Salad")
        for _, value := range slice {
		fmt.Println("Food: ",value)
	}		
}