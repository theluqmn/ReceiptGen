package main

import (
	"bufio"
	"fmt"
	"os"
	"container/list"
	"strings"
)

type Item struct {
	name string
	price float64
}

func main() {
	fmt.Println("Hello, World!")

	// Get user input
	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Greet user
	fmt.Printf("Hello, %s!\n", name)

	itemList := list.New()

	// Tax rate
	fmt.Println("Enter tax rate: (%)")
	var taxRate float64
	fmt.Scanln(&taxRate)
	taxRate /= 100

	// Items to add
	fmt.Println("Amount of items to add:")
	var amount int
	fmt.Scanln(&amount)

	for i := 0; i < amount; i++ {
		// Get item data
		fmt.Println("- Enter item name:")
		itemName, _ := reader.ReadString('\n')
		itemName = strings.TrimSpace(itemName)

		fmt.Println("- Enter item price:")
		var itemPrice float64
		fmt.Scanln(&itemPrice)

		// Add item to list
		itemList.PushBack(Item{name: itemName, price: itemPrice})
	}

	var totalValue float64
	var totalTax float64
	var totalDue float64

	// Print list
	fmt.Println("------------------------------------------------")
	fmt.Println("List:")
	for e := itemList.Front(); e != nil; e = e.Next() {
		item := e.Value.(Item)
		fmt.Printf("%s: %.2f\n", item.name, item.price)

		totalValue += item.price
		totalTax += (item.price * taxRate)
	}

	totalDue = totalValue + totalTax

	fmt.Println("------------------------------------------------")
	fmt.Printf("Total value: %.2f\n", totalValue)
	fmt.Printf("Total tax: %.2f\n", totalTax)
	fmt.Printf("Total due: %.2f\n", totalDue)
	fmt.Println("------------------------------------------------")
}
