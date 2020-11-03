package main
import "fmt"

func main() {
	fmt.Println("We are gonna ")//Printing the day of week using input number
	fmt.Println("Enter the day number like 1 2 3..so on")
	//var day int
	//fmt.Scanf("%d",&day)
	switch day:=4;day{
	case 1:
	fmt.Println("Today is monday")
	case 2:
	fmt.Println("Today is Tuesday")
	case 3:
	fmt.Println("Today is wendsday")
	case 4:
	fmt.Println("Today is Thurday")
	case 5:
	fmt.Println("Today is Friday")
	case 6:
	fmt.Println("Today is Saturday")
	case 7:
	fmt.Println("Today is sunday")
	default:
	fmt.Println("Invalid day")
	}
}
