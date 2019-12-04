package main

import (
	"fmt"
)

func bothChannel(ch chan int) {
	ch <- 6
	fmt.Println(<-ch)
}

func sendChannel(ch chan<- int) {
	ch <- 555
}

func recvChannel(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	// 채널 생성
	ch := make(chan int)

	go func() {
		// defer wait.Done()
		ch <- 999
	}()

	// send가 있을때까지 기다린다.
	i := <-ch
	fmt.Println("i: ", i)

	// ch2 := make(chan int)
	// ch2 <- 888 // 수신하는 goroutine이 없어서 deadlock
	// i2 := <-ch2 // 송신하는 goroutine이 없어서 deadlock
	// fmt.Println(i2)

	ch3 := make(chan int, 1)

	// defer close(ch3)

	ch3 <- 777
	fmt.Println(<-ch3)

	bothChannel(ch3)

	sendChannel(ch3)
	recvChannel(ch3)

	ch3 <- 44
	close(ch3)
	println(<-ch3)

	if _, success := <-ch3; !success {
		println("empty channel")
	}

	ch4 := make(chan int, 2)
	ch4 <- 1
	ch4 <- 2
	close(ch4)

	for {
		if i, success := <-ch4; success {
			println(i)
		} else {
			break
		}
	}

	ch5 := make(chan int, 2)
	ch5 <- 1
	ch5 <- 2
	close(ch5)

	for i := range ch5 {
		println(i)
	}

	ch6 := make(chan bool)
	ch7 := make(chan bool)

	go runChannel6(ch6)
	go runChannel7(ch7)

EXIT:
	for {
		select {
		case r := <-ch6:
			println(r)
		case r := <-ch7:
			println(r)
			// 그냥 break를 하면 select deadlock이 발생되어 이 방식을 써야한다.
			// 아직까지 정확한 이유는 못찾았다.
			break EXIT
			// 단순 break를 걸게되면 default가 무한하게 호출된다.
			// break
			// default:
			// 	println("select default")
			// 	break
		}
	}
}

func runChannel6(ch chan<- bool) {
	ch <- true
}

func runChannel7(ch chan<- bool) {
	ch <- true
}
