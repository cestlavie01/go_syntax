package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("1.txt")
	fmt.Println("Done") // 이 문장 실행됨
}

func openFile(fn string) {
	defer func() {
		fmt.Println("defer: before recover")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	defer func() {
		fmt.Println("defer: before open 1")
	}()

	defer func() {
		fmt.Println("defer: before open 2")
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	// 출력 안됨
	defer func() {
		fmt.Println("defer: after open")
	}()

	defer f.Close()
}
