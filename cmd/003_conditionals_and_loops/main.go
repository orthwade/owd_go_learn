package main

import (
	"fmt"
	"math/rand"
)

func tryDefer() {
	defer fmt.Println("This will be printed at the end of the function")
	fmt.Println("This will be printed first")
}

func main() {
	var arrStr = [3]string{"a", "b", "c"}

	fmt.Println("If else example. Checking if str is equal to a, b or c")

	randomStr := arrStr[rand.Intn(len(arrStr))]

	if randomStr == "a" {
		fmt.Println("str is a")
	} else if randomStr == "b" {
		fmt.Println("str is b")
	} else if randomStr == "c" {
		fmt.Println("str is c")
	} else {
		fmt.Println("Error. str is not a, b or c")
	}

	fmt.Println("Switch example. Checking if str is equal to a, b or c")

	switch randomStr {
	case "a":
		fmt.Println("str is a")
	case "b":
		fmt.Println("str is b")
	case "c":
		fmt.Println("str is c")
	default:
		fmt.Println("Error. str is not a, b or c")
	}

	fmt.Println("For loop example. Looping from 0 to 10")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("For loop example. Looping from 0 to 10 with continue")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Skipping 5")
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("For loop example. Looping from 0 to 10 with break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking at 5")
			break
		}
		fmt.Println(i)
	}

	fmt.Println("For loop example. Looping from 0 to 10 with x declared outside")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("For loop example. Looping through array")
	for i, v := range arrStr {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println("For loop example. Looping through array without index")
	for _, v := range arrStr {
		fmt.Printf("Value: %s\n", v)
	}

	fmt.Println("For loop example. Looping through array indices")
	for i, _ := range arrStr {
		fmt.Printf("Index: %d\n", i)
	}

	fmt.Println("Defer example")

	tryDefer()

	fmt.Println("After defer")

}
