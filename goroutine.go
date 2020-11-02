package main

import ("fmt";"sync")//This is the program using goroutine in synchronisation
func foo(wg *sync.WaitGroup){
	fmt.Println("This is new goroutine");
	wg.Done()
}
func main() {
 	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("main go routine")
}
