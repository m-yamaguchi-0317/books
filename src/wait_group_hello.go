package main

import (
	"fmt"
	"sync"
)

func main() {
	helloListPreference()
}

// ゴルーチンの同期
func hello() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello!")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait() //合流ポイント、mainの終了をブロック
}

// ゴルーチン間のアドレス空間の共有
func helloTwo() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation) // welcomeが出力される -> ゴルーチンは生成元のゴルーチンとアドレス空間を共有するということ

}

// 生成元とゴルーチンのアドレス空間の共有の落とし穴
func helloList() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello!", "greeting!", "good day!"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	// "good day!"が3回出力される
	// -> ゴルーチンが生成される前にfor文が終了してしまうので、３つにゴルーチンが全て、"good dayを出力してしまう"
}

// 生成元とゴルーチンのアドレス空間の共有への対処
func helloListCopy() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello!", "greeting!", "good day!"} {
		wg.Add(1)
		go func(salutation string) { //関数に引数を指定してあげる
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // salutatinのコピーを渡してあげる
	}
	wg.Wait()
	// "hello!", "greeting!", "good day!"の順に出力される
	// -> ゴルーチン（無名関数）の生成の際に、salutationのコピーを渡しているので、適切に表示されるy
}

// 生成元とゴルーチンのアドレス空間の共有の確認（参照型）
func helloListPreference() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello!", "greeting!", "good day!"} {
		wg.Add(1)
		go func(salutation *string) { //関数に引数を指定してあげる
			defer wg.Done()
			fmt.Println(*salutation)
		}(&salutation) // 現在の繰り返しのstringオブジェクトをクロージャーに渡してあげる
	}
	wg.Wait()
	// "good day!"の順に出力される
	// -> ゴルーチン（無名関数）の生成の際に、salutationの参照型を渡しているので、helloList()と同じ結果になる
}
