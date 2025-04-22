package main

import "fmt"

func main() {
	arrInt := [3]int{1, 2, 3}

	sliceInt := []int{1, 2, 3}

	fmt.Println("Array:", arrInt)
	fmt.Println("Slice:", sliceInt)

	fmt.Println("Adding 4 to slice")
	sliceInt = append(sliceInt, 4)
	fmt.Println("Slice:", sliceInt)

}
