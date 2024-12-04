package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to the bank.")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance.")
		fmt.Println("2. Withdraw money.")
		fmt.Println("3. Deposit money.")
		fmt.Println("4. Exit. Bye.")

		var choice int

		fmt.Println("Please enter your choice (number):")
		fmt.Scan(&choice)

		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			logBalance(accountBalance)
			continue
		case 2:
			withdraw(&accountBalance)
			continue
		case 3:
			deposit(&accountBalance)
			continue
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return
			// break  // break would not escape the for loop, but only switch.
		}
	}

}

func logBalance(accountBalance float64) {
	fmt.Println("Your balance is: ", accountBalance)
}

func withdraw(accountBalance *float64) {
	fmt.Println("Enter your withdrawal amount: ")
	var withdrawal float64
	fmt.Scan(&withdrawal)

	if withdrawal <= 0 {
		fmt.Println("Withdrawal amount cannot be negative or 0.")
	} else if withdrawal > *accountBalance {
		fmt.Println("Withdrawal amount cannot be lareger than your account balance.")
		logBalance(*accountBalance)
	} else {
		*accountBalance -= withdrawal
		fmt.Println("Balance updated after withdrawal. You new balance is: ", *accountBalance)
	}
}

func deposit(accountBalance *float64) {
	var deposit float64

	fmt.Println("Enter your deposit: ")
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("Deposit amount cannot be negative or 0.")
	} else {
		*accountBalance += deposit
		fmt.Println("Balance updated after deposit. You new balance is: ", *accountBalance)
	}
}
