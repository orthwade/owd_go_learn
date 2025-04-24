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

	mapStrInt := make(map[string]int)

	mapStrInt["one"] = 1
	mapStrInt["two"] = 2

	fmt.Println("Map:", mapStrInt)

	sliceIntCopy := make([]int, len(sliceInt))
	copy(sliceIntCopy, sliceInt)
	fmt.Println("Slice Copy:", sliceIntCopy)

	val, ok := mapStrInt["one"]
	if ok {
		fmt.Println("Value for 'one':", val)
	} else {
		fmt.Println("'one' not found in map")
	}

	val, ok = mapStrInt["three"]
	if ok {
		fmt.Println("Value for 'three':", val)
	} else {
		fmt.Println("'three' not found in map")
	}

	delete(mapStrInt, "two")
	fmt.Println("Map after deletion:", mapStrInt)

}
