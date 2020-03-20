package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	measureGoroutine()
}

func measureGoroutine() {
	memCousumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // 計測のためにゴルーチンをメモリにとどめたいので、終了しないゴルーチンを生成

	const numGoroutines = 1e4 // ゴルーチの数を指定
	wg.Add(numGoroutines)
	before := memCousumed() //ゴルーチンの生成前のメモリ消費量を計測
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memCousumed() // ゴルーチン生成後のメモリ消費量を計測
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

// ゴルーチン間でのメッセージの共有
func sendMessageWhileGoroutine() {

}
