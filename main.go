package main

import (
	"bufio"
	"fmt"
	"os"
	"container/list"
	"strings"
	"time"
	"os/exec"
	"strconv"
)

type Item struct {
	name string
	price float64
}

func main() {
	// Clear console
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("Hello, World!")

	// Get user input
	fmt.Printf("- Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	itemList := list.New()

	// Tax rate
	fmt.Printf("- Enter tax rate percentage: ")
	taxRateStr, _ := reader.ReadString('\n')
	taxRateStr = strings.TrimSpace(taxRateStr)
	taxRate, err:= strconv.ParseFloat(taxRateStr, 64)
	if err != nil {
		fmt.Println("Invalid tax rate")
		return
	}
	taxRate /= 100

	// Items to add
	fmt.Printf("- Amount of items to add: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	fmt.Println("\n------------------------------------------------")
	fmt.Println("Now, enter items.")

	for i := 0; i < amount; i++ {
		// Get item data
		fmt.Printf("%d - Enter item name: ", i + 1)
		itemName, _ := reader.ReadString('\n')
		itemName = strings.TrimSpace(itemName)

		fmt.Printf("%d - Enter item price: $", i + 1)
		var itemPrice float64
		fmt.Scanln(&itemPrice)

		// Add item to list
		itemList.PushBack(Item{name: itemName, price: itemPrice})
	}

	var totalValue float64
	var totalTax float64
	var totalDue float64

	// Print receipt
	fmt.Println("\n------------------------------------------------")
	fmt.Println("Receipt for:", name)
	fmt.Println("Issued at:", time.Now().Format("02-01-2006 15:04:05"))
	fmt.Println("------------------------------------------------")
	fmt.Printf("Items: %v\n", itemList.Len())

	var count int = 1
	for e := itemList.Front(); e != nil; e = e.Next() {
		item := e.Value.(Item)
		fmt.Printf("%v/%v - '%s': $%.2f\n", count, itemList.Len(), item.name, item.price)

		totalValue += item.price
		totalTax += (item.price * taxRate)
		count++
	}

	totalDue = totalValue + totalTax
	taxRatePercentage := taxRate * 100

	fmt.Println("------------------------------------------------")
	fmt.Printf("Total value: $%.2f\n", totalValue)
	fmt.Printf("Total tax: $%.2f (%v%%)\n", totalTax, taxRatePercentage)
	fmt.Printf("Total due: $%.2f\n", totalDue)
	fmt.Println("------------------------------------------------")
}