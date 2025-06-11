package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func queue(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	randomInt := rand.Intn(3-1+1)+1
	
	fmt.Println("Queue:", msg, "takes", randomInt, "seconds")
	time.Sleep(time.Duration(randomInt) * time.Second)
}

func main() {
	for {
	randomNumber := rand.Intn(3-1+1)+1
	
	var wg sync.WaitGroup
	fmt.Printf("----Now you have: %d queues----\n", randomNumber)
	for q := 1; q <= randomNumber; q++ {
		wg.Add(1)
		
		go queue(&wg, fmt.Sprintf("%d", q))
		
		wg.Wait()
	}
	}
}
