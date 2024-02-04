package main

import (
	"fmt"
	"sync"
	"time"
)

func syncFunction() {
	for i := 0; i < 3; i++ {
		fmt.Println("syncFunction, i:", i)
	}
}

func printValue(i int) {
	fmt.Println("printValue, i:", i)
}

func asyncFunction() {
	for i := 0; i < 3; i++ {
		fmt.Println("creating go routine for, i:", i)
		go printValue(i)
		// defer wg.Done()
	}
}

func printValueWg(i int, wg *sync.WaitGroup) {
	fmt.Println("printValueWg, i:", i)
	defer wg.Done()
}

func asyncFunctionWg(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println("asyncFunctionWg, creating go routine for, i:", i)
		go printValueWg(i, wg)
		// defer wg.Done()
	}
}

func main() {
	fmt.Println("go routines example")

	var wg sync.WaitGroup
	wg.Add(3)
	go asyncFunctionWg(&wg)
	go asyncFunction()
	time.Sleep(10 * time.Microsecond)
	syncFunction()
	wg.Wait()
}
