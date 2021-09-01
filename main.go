package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// throw function inside
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	// var nameOne string = "hello"
	// var nameTwo = "test"

	// // This equal to ''
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// // Can use only in function
	// nameFour := "yoshi"

	// fmt.Println((nameFour))

	// // Number , int, float
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// // // bit & memory
	// // var num1 int8 = 25
	// // var numThree uint8 = 255
	// fmt.Println(ageOne, ageTwo, ageThree)

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 88888855646.7
	// fmt.Println(scoreOne, scoreTwo)

	// // Print
	// fmt.Println("hello, ", 35)

	// //Printf
	// fmt.Printf("my age is %v and my name is %v \n", 23, "benning")
	// fmt.Printf("my age is %0.2f and my name is %v \n", 23.888, "benning")

	// // SprintF
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", 23, "benning")
	// fmt.Println((str))

	// fixed length
	var ages [3]int = [3]int{20, 25, 30}

	var ages2 [3]int = [3]int{20, 25, 30}

	fmt.Println(ages, len(ages), ages2)

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 25

	scores = append(scores, 85)

	var scores2 = scores[:2]
	fmt.Println(scores, len(scores), scores2)

	fmt.Println(strings.ToUpper("fuck"))

	fmt.Println(strings.Split("I Love You", " "))

	// loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range scores {
		fmt.Printf("The position at index %v the value is %v \n", index, value)
	}

	if scores[2] <= 25 {
		fmt.Println("test")
	}

	// sayGreeting("Benning")
	// sayBye("Benning")
	cycleNames([]string{"Benning", "Muse", "Hello"}, sayGreeting)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f circle 2 is %0.3f \n", a1, a2)
}
