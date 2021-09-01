package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func propmptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add items, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			propmptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		propmptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			propmptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		propmptOptions(b)
	case "s":
		b.save()
		fmt.Println("you choose to save the bill", b.name)
	default:
		fmt.Println("That was not a valid option...")
		propmptOptions(b)
	}
}

func main() {
	// sayHello("Benning")
	// showScore()
	// mapValue()

	// pointerExample()

	// myBill := newBill("Benning")
	// myBill.addItem("Burger", 15)
	// myBill.updateTip(30)
	// fmt.Println(myBill.format())

	mybill := createBill()
	propmptOptions(mybill)
}
