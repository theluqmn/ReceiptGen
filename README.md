# RECEIPTGEN

> [!WARNING]
> This is my first project made using Go.

## Purpose

Returns a receipt-like output on the terminal, which includes all the items you added, alongside its price and tax. The total is calculated at the bottom.

## How it works

If you want the .exe, its on [Releases](https://github.com/luq-mn/ReceiptGen/releases).

1. Enter your name
2. Enter the tax rate (0-100)%
3. Enter the amount of items that will be in this list
4. The list will loop until the amount of items in list meets the amount you set previously.
    * Enter the item name
    * Enter the item price
5. Enter your payment method. Choose between:
    * Cash
    * Credit Card
    * Debit Card
    * E-Wallet
6. The full list will be outputted to the terminal, formatted like a receipt. Includes the total due, total price, and tax.

## To-do

* [x] it works
* [x] payment method
* [ ] input validation
* [x] output with more details
* [ ] output as .txt
