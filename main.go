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

func clear() {
	// Clear console
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	fmt.Println("Hello, World!")

	fmt.Println("Luq's ReceiptGen - Receipt Calculator/Generator")

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

	clear()
	fmt.Println("Now, enter items.")

	for i := 0; i < amount; i++ {
		// Get item data
		fmt.Printf("%d/%v - Enter item name: ", i + 1, amount)
		itemName, _ := reader.ReadString('\n')
		itemName = strings.TrimSpace(itemName)

		fmt.Printf("%d/%v - Enter item price: $", i + 1, amount)
		var itemPrice float64
		fmt.Scanln(&itemPrice)

		// Add item to list
		itemList.PushBack(Item{name: itemName, price: itemPrice})
		clear()
	}

	// Ask for payment method

	var paymentMethod string

	clear()
	fmt.Println("After entering your payment method,\nthe receipt will be outputted to the console.")
	fmt.Println("Options:\n- [1] Cash\n- [2] Credit Card\n- [3] Debit Card\n- [4] E-Wallet")

	fmt.Printf("- Enter payment method [1-4]: ")
	paymentMethod, _ = reader.ReadString('\n')
	paymentMethod = strings.TrimSpace(paymentMethod)

	paymentMethodInt, err := strconv.Atoi(paymentMethod)
	if err != nil {
		fmt.Println("Invalid payment method")
		return
	}

	if paymentMethodInt == 1 {
		paymentMethod = "Cash"
	} else if paymentMethodInt == 2 {
		paymentMethod = "Credit Card"
	} else if paymentMethodInt == 3 {
		paymentMethod = "Debit Card"
	} else if paymentMethodInt == 4 {
		paymentMethod = "E-Wallet"
	} else {
		fmt.Println("Invalid payment method")
		return
	}

	var totalValue float64
	var totalTax float64
	var totalDue float64

	// Print receipt
	fmt.Println("\n------------------------------------------------")
	fmt.Println("Receipt for:", name)
	fmt.Println("Issued at:", time.Now().Format("02-01-2006 15:04:05"))
	fmt.Println("------------------------------------------------")
	fmt.Printf("Items: %v\n\n", itemList.Len())

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
	fmt.Printf("Total tax: $%.2f (%.2f%%)\n", totalTax, taxRatePercentage)
	fmt.Printf("\nTotal due: $%.2f\n", totalDue)
	fmt.Printf("Payment method: %s\n", paymentMethod)
	fmt.Println("------------------------------------------------")
}