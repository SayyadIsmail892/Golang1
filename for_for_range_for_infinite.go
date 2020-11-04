package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a program for different form of loops in go")
	for i:=0;i<=5;i++{//This is one form of for loop
	fmt.Println("This is sayyad ismail from kloudone")
	}
	i:=0
	for i<=8{//this is similar to while loop
	fmt.Println("This is another form of for loop like subtitute to while loop in other programming languages")
	i++
	}
	fruits:=[]string{"apple","banana","pineapple","orange"}
	for i,j:=range fruits{
	fmt.Println(i,j)
	}
	for{  //This is infinite loop
	fmt.Println("hai infinte loop")
	break//just i am ending it but it is infinite loop
	}
	
}
