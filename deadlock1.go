package main

import (
	"fmt"
	"sync")
func dostuff(c1 chan int,c2 chan int){  //This is the program for deadlock
<-c1
c2<-1
wg.Done()
}
var wg sync.WaitGroup  
func main() {
	//var wg sync.WaitGroup  
	//Deadlock occurs due to circular dependencies between goroutines
	fmt.Println("This is deadlock example")
	wg.Add(2)
	//fmt.Println("This is deadlock example")
	ch1:=make(chan int)
	ch2:=make(chan int)
	go dostuff(ch1,ch2)
	go dostuff(ch2,ch1)
	wg.Wait()
}
