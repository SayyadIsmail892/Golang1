package main

import "fmt"

func main()   {
	fmt.Println("Hello, playground")
	var age int
	age=25
	//This is if example
	fmt.Println("The if statemet example")
	if (age>=25) {
	fmt.Println("He is adult")
	fmt.Println("This is if statement")
	}
	//This is if else example
	fmt.Println("The if else statement example")
	if age<18{
	fmt.Println("He is child")
	}else if age>=18{
	fmt.Println("He is adult")
	}else{
	  fmt.Println("He is may be different")
	}
	
}
