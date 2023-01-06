package main
import "fmt"

func main(){
	switch day :=4; day {
		case (1): fmt.Printf("Monday")
		case (2): fmt.Printf("Tuesday")
		case (3): fmt.Printf("Wednesday")
		case (4): fmt.Printf("Thursday")
		default: fmt.Printf("Invalid")
	}
}