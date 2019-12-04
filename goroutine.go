package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")

	time.Sleep(time.Second * 1)

	// wait이 없기 때문에 출력이 정상적이지 않다.
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("Async3 ***", i)
		}
	}()

	var wait sync.WaitGroup
	wait.Add(1)

	go func(name string) {
		defer wait.Done()

		for i := 0; i < 2; i++ {
			fmt.Println("Async4 ***", name, i)
		}
	}("cobus")

	// 여기에서 wait을 하기 때문에 Async4가 모두 출력된다.
	// 하지만 여전히 Async3은 온전한 실행 완료를 100% 보장받지 못한다.
	wait.Wait()
}
