package main

import "fmt"
import "sync" //This is the program two goroutines sharing same variable
func inc(){
	mut.Lock()
	i=i+1
	mut.Unlock()
	wg.Done()
	}
var i int=0
var wg sync.WaitGroup
var mut sync.Mutex
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
