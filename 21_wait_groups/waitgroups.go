package main

import (
	"fmt"
	"sync"
)

func task(num int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(num)
}

func main() {

	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
