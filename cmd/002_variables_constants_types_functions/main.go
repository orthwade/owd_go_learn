package main

import "fmt"

// Variables
var name string = "Orthwade"
var age int = 30
var ageFloat float32 = float32(age)
var height float32 = 1.80

// Formatted message
var message string = fmt.Sprintf(
	"My name is %s, I am %.2f years old and my height is %.2f meters.",
	name, ageFloat, height,
)

func sum(args ...int) int {
	// Sum all the integers in args
	total := 0

	for _, arg := range args {
		total += arg
	}

	return total
}

func add(v_0 int, v_1 int) int {
	return v_0 + v_1
}

func applyFunc(f func(int, int) int, val_1 int, val_2 int) int {
	return f(val_1, val_2)
}

func showLongestName(sliceName []string) (string, int) {
	// Find the longest name in the slice
	longestName := ""
	longestLength := 0

	for _, name := range sliceName {
		if len(name) > longestLength {
			longestLength = len(name)
			longestName = name
		}
	}

	return longestName, longestLength
}

func main() {
	// Print the message
	fmt.Println(message)

	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))

	fmt.Println("Sum of 10, 20:", applyFunc(add, 10, 20))

	lambdaAdd := func(v_0 int, v_1 int) int {
		return v_0 + v_1
	}

	fmt.Println("Lambda sum of 10, 20:", lambdaAdd(10, 20))

	name, i := showLongestName([]string{"John", "Jane", "Alexander", "Sophia"})

	fmt.Println("Longest name:", name, "with length:", i)

	// Basic function (placeholder for additional code)
}
