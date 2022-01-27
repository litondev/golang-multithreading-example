package main

import (
	"fmt"
	// "time"
	"sync"
)

var wg sync.WaitGroup

func insertData(thread int){
	fmt.Println(thread);
	wg.Done();

}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go insertData(i)
	}

	wg.Wait()
}
