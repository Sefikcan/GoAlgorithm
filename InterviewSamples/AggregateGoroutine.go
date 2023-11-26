package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	userName := fetchUser()

	respCh := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2) // because of we have 2 async operation

	// we waits for the longest operation
	go fetchUserLikes(userName, respCh, wg)
	go fetchUserMatch(userName, respCh, wg)

	wg.Wait() // block until 2 wg.Done()

	close(respCh) // if we don't close the channel we will get a deadlock error

	for resp := range respCh {
		fmt.Println("resp:", resp)
	}

	//fmt.Println("likes:", likes)
	//fmt.Println("match:", match)

	fmt.Println("took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Şefik"
}

func fetchUserLikes(userName string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respCh <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respCh <- "Yagmur"
	wg.Done()
}
