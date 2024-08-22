package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

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


func storeInDB(name string, total float64, items *list.List) {
	db, err := sql.Open("sqlite3", "receipt.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS receipts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		total FLOAT,
		timestamp TEXT
	)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		receipt_id INTEGER,
		item_name TEXT,
		item_price FLOAT,
		FOREIGN KEY(receipt_id) REFERENCES receipts(id)
	)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Insert into receipts table
	result, err := tx.Exec(`INSERT INTO receipts (name, total, timestamp) VALUES (?, ?, ?)`,
		name, total, time.Now().Format("02-01-2006 15:04:05"))
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	receiptID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	// Insert items into items table
	for e := items.Front(); e != nil; e = e.Next() {
		item := e.Value.(Item)
		_, err = tx.Exec(`INSERT INTO items (receipt_id, item_name, item_price) VALUES (?, ?, ?)`,
			receiptID, item.name, item.price)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	clear()
	fmt.Println("we gonna cook some reciepts up baby")

	fmt.Println("Luq's ReceiptGen - Receipt Calculator/Generator")

	// Get user input
	fmt.Printf("(1/6) - Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	itemList := list.New()

	// Tax rate
	fmt.Printf("(2/6) - Enter tax rate: ")
	taxRateStr, _ := reader.ReadString('\n')
	taxRateStr = strings.TrimSpace(taxRateStr)
	taxRate, err:= strconv.ParseFloat(taxRateStr, 64)
	if err != nil {
		fmt.Println("Invalid rate")
		return
	}
	taxRate /= 100

	// Items to add
	fmt.Printf("(3/6) - Amount of items to add: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	clear()
	fmt.Println("(4/6) - Now, item names.")

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
	fmt.Println("After using your payment method,\nthe receipt will be outputted to the console.")
	fmt.Println("Options:\n- [1] Cash\n- [2] Card\n- [3] E-Wallet\n- [4] Pay later with Afterpay")

	fmt.Printf("(5/6) - Enter payment method [1-4]: ")
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
		paymentMethod = "Card"
	} else if paymentMethodInt == 3 {
		paymentMethod = "E-Wallet"
	} else if paymentMethodInt == 4 {
		paymentMethod = "Afterpay"
	} else if paymentMethodInt == 5 {
		paymentMethod = "You screwed up"
	} else {
		fmt.Println("Invalid payment method")
		return
	}

	// Discount
	var discount float64

	clear()
	fmt.Println("Enter discount percentage. Discount is applied before taxes.\nLeave blank for no discount.")
	fmt.Printf("(6/6) - Enter discount percentage: ")
	discountStr, _ := reader.ReadString('\n')
	discountStr = strings.TrimSpace(discountStr)

	if  discountStr == "" {
		discount = 0
	} else {
		discount, err = strconv.ParseFloat(discountStr, 64)
		if err != nil {
			fmt.Println("Invalid discount")
			return
		}

		discount /= 100
	}

	
	var totalValue float64
	var totalTax float64
	var totalDue float64

	// Print receipt
	clear()
	fmt.Println("\n------------------------------------------------")
	fmt.Println("Receipt issued to:", name)
	fmt.Println("On:", time.Now().Format("02-01-2006 15:04:05"))
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

	if discount != 0 {
		totalDue = totalValue - (totalValue * discount)
	} else {
		totalDue = totalValue
	}

	totalDue += totalTax
	taxRatePercentage := taxRate * 100

	storeInDB(name, totalDue, itemList)

	fmt.Println("------------------------------------------------")
	fmt.Printf("Subtotal: $%.2f\n", totalValue)
	fmt.Printf("\nDiscount: $%.2f (%.2f%%)\n", discount * totalValue, discount * 100)
	fmt.Printf("Tax: $%.2f (%.2f%%)\n", totalTax, taxRatePercentage)
	fmt.Printf("\nTotal: $%.2f\n", totalDue)
	fmt.Printf("Payment method: %s\n", paymentMethod)
	fmt.Println("------------------------------------------------")
}
