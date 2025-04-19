package main

import "fmt"

var name string = "Orthwade"

var age int = 30

var ageFltoat float32 = float32(age)

var height float32 = 1.80

var message string = fmt.Sprintf("My name is %s, I am %f years old and my height is %.2f meters.", name, ageFltoat, height)

func main() {
	fmt.Println(message)
}
