package main

import "fmt"
import "sync" //This is the program two goroutines sharing same variable
func inc(){
	i=i+1
	wg.Done()
	}
var i int=0
var wg sync.WaitGroup
func main() {
	//var i int=0
	//var wg sync.WaitGroup
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println("The value of i is now",i)
	fmt.Println("Hello, playground")
}
