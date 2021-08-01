package main

import (
	"fmt"
	"math"
)

// to run an app, exec "go run main.go" from terminal. You should go to this folder in the terminal

// getCoins is a simple function that use price, and coin nominal as input args,
// and return the amount of coins and remaining price as a result
func getCoins(price, coin float64) (int64, float64) {
	amountOfCoins := int64(price / coin) // I use int64 to get only "integer" part here
	price = math.Mod(price, coin)        // https://golang.org/src/math/mod.go
	price = math.Round(price*100) / 100  // https://www.geeksforgeeks.org/math-round-function-in-golang-with-examples/

	return amountOfCoins, price // yes, in go we can return a few values from a function
}

func main() {
	// solution #1
	// ____________________________________________
	var price float64                    // "price" is just a box where I can put some values
	price = 1.15                         // "=" means variable assignment
	_50s, price := getCoins(price, 0.50) // _50s: I use "_" 'cause we can't start a variable name from a number. Usually it's not the best practice :)
	_25s, price := getCoins(price, 0.25) // 0.50, 0.10 etc is a good case to use const in go.
	_10s, price := getCoins(price, 0.10) // https://blog.golang.org/constants <- interesting article for reading but not for beginners 												  but looks like
	_5s, price := getCoins(price, 0.05)
	_1s, _ := getCoins(price, 0.01)

	fmt.Printf("50 cents: %d coins\n25 cents: %d coins\n10 cents: %d coins\n5 cents: %d\n1 cent: %d coins\n",
		_50s, _25s, _10s, _5s, _1s) // https://gobyexample.com/string-formatting <- more info about string formatting. \n means "end of the line"
	// ____________________________________________

	// solution #2 (just comment out previous solution and uncomment this to play with solution #2
	// ____________________________________________
	// var price float64 // "price" is just a box where I can put some values
	// price = 1.15      // "=" means variable assignment

	// for price > 0 {
	// 	price = math.Round(price*100) / 100
	// 	switch { // https://golangdocs.com/switch-statement-in-golang
	// 	case price >= 0.50:
	// 		fmt.Println("50 cents coin!")
	// 		price = price - 0.50
	// 	case price >= 0.25:
	// 		fmt.Println("25 cents coin!")
	// 		price = price - 0.25
	// 	case price >= 0.10:
	// 		fmt.Println("10 cents coin!")
	// 		price = price - 0.10
	// 	case price >= 0.05:
	// 		fmt.Println("5 cents coin!")
	// 		price = price - 0.05
	// 	case price >= 0.01:
	// 		fmt.Println("1 cent coin!")
	// 		price = price - 0.01
	// 	}
	// }
	// ____________________________________________
}
