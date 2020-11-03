package main

import (
	"fmt"
	"sync"
)
var on sync.Once   /*This is the example program for syncronization in Go using Goroutines
                     here we used sync.Once package for initalization purpose*/
                     
func setup(){
fmt.Println("Init")
}
func dostuff(){
on.Do(setup)
fmt.Println("Hello")
wg.Done()
}
var wg sync.WaitGroup
func main() {
	fmt.Println("Hello, playground")
	//var wg sync.WaitGroup
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}