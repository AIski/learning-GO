package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"pl.aski/bank/fileOps"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Print("ERROR: ", err)
		fmt.Print("---------")
		// panic("Calling panic function!")
	}

	fmt.Println("Welcome to the bank.")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int

		fmt.Println("Please enter your choice (number):")
		fmt.Scan(&choice)

		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			logBalance(accountBalance)
			fileOps.SaveFloatToFile(accountBalanceFile, accountBalance)
			continue
		case 2:
			withdraw(&accountBalance)
			fileOps.SaveFloatToFile(accountBalanceFile, accountBalance)
			continue
		case 3:
			deposit(&accountBalance)
			fileOps.SaveFloatToFile(accountBalanceFile, accountBalance)
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
