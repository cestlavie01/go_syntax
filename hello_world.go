package main

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}
}

func say(msg ...string) {
	for _, word := range msg { // _ means I don't want to use index
		println(word)
	}

	for index, _ := range msg {
		println(index)
	}
}

func printWithPointer(name *string) {
	println(*name)
	*name = "cobus"
}

func classicSum(nums ...int) int {
	r := 0
	for _, num := range nums {
		r += num
	}
	return r
}

// func can return more than 2 values
func goSum1(nums ...int) (int, int) {
	count := 0
	r := 0
	for _, num := range nums {
		count++
		r += num
	}
	return r, count
}

// it can define return parameter's name
func goSum2(nums ...int) (total int, count int) {
	for _, num := range nums {
		count++
		total += num
	}
	return // must to be return
}

type caculator func(int, int) int

func calc(f caculator, a int, b int) int {
	return f(a, b)
}

// This function looks like return 2 variables but it will return only function.
// annonymous function doesn't include "i" but you can see value's still countinue.
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var f float32 = 123.5
	var i int = int(f)
	var f2 float32 = float32(i)
	println(f, f2, i)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(str, str2, bytes)

	checkType(f)   // unknown
	checkType(i)   // int
	checkType(str) // string

	members := []string{"jtlee", "cobus", "ddukddak"}
	for index, member := range members {
		println(index, member)
	}
	/*
		0 jtlee
		1 cobus
		2 ddukddak
	*/

	i = 0

L1:
	for { // infinit loop
		if i == 0 {
			break L1 // after break, for loop will be ignored
			// break L2 // error: because label L2 doesn't define front of this line.
		}

		println("inner loop") // never print this
	}
	// L2: // error: because this label is never called

	println("end of loop")

	say("hello", "world", "!!")

	name := "jtlee"
	printWithPointer(&name)
	println(name)

	println(classicSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	total, count := goSum1(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(total, count)

	total, count = goSum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(total, count)

	// annoymous function
	sum := func(i int, j int) (total int) {
		total = i + j
		return
	}

	println(sum(1, 2))

	// using callback function even if annoymous function
	println(calc(sum, 1, 2))

	// Closure
	println("----------- closure")
	// next, iNext := nextValue() // assignment mismatch: 2 variables but nextValue returns 1 values
	next := nextValue()
	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1
	println(anotherNext()) // 2
	println(anotherNext()) // 3

	// Array
	println("---------- array")
	var a [3]int
	a[0] = 1 // zero based index
	a[1] = 2
	a[1] = 3
	println("a[0]: ", a[0])

	var b = [3]int{1, 2, 3}
	println("b[2]: ", b[2])

	var c = [...]int{1, 2, 3, 4, 5, 6}
	println("c[5]: ", c[5])

	var d [2][3]int
	d[0][1] = 19
	println("d[0][0]: ", d[0][0])
	println("d[0][1]: ", d[0][1])

	var e = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // have to put comma
	}
	println("e[0][2]: ", e[0][2])
	println("e[1][2]: ", e[1][2])

	// Slice -- http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice
	// it's like slice of Python but HAVE TO KNOW about managing of memory
}
