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

func main() {
	// Print the message
	fmt.Println(message)

	// Basic function (placeholder for additional code)
}
