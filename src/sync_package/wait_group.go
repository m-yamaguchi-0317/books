package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	waitTest()
}

func waitTest(){
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1*time.Second)
	}()

	wg.Add(1) //wg.Add()は監視対象のゴルーチンの生成前に定義するのが慣習
	go func(){
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2*time.Second)
	}()

	wg.Wait()
	fmt.Printf("Const time is %v\n", time.Since(now))
	fmt.Println("All Goroutine comprite")
}