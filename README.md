# RECEIPTGEN

> [!WARNING]
> This is my first project with Go. I am planning on sticking with it as my main go-to compiled language.

## Purpose

Returns a receipt-like output on the terminal, which includes all the items you added, alongside its price and tax. The total is calculated at the bottom.

## How it works

If you want the .exe, its on Releases.

1. Enter your name
2. Enter the tax rate (0-100)%
3. Enter the amount of items that will be in this list
4. The list will loop until the amount of items in list meets the amount you set previously.
    * Enter the item name
    * Enter the item price
5. The full list will be returned, with the total value, tax, and due.

## To-do

* [x] it works
* [ ] input validation
* [ ] output with more details
* [ ] output as .txt
