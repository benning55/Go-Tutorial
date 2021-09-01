package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

func showScore() {
	fmt.Printf("You shit %0.2f \n", 45.5)
}

func mapValue() {
	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
}

// Get address of variable to update
func updateName(m *string) {
	*m = "Benning is more awsome"
}

func pointerExample() {
	name := "tifa"

	// show address of variable name
	fmt.Println(&name)

	m := &name
	fmt.Println(*m)
	updateName(m)
	fmt.Println(*m)

}
