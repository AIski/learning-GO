package main

import (
	"errors"
	"fmt"
	"os"
)

const accountBalanceFile = "balance.txt"

func main() {

	revenue, err := getUserInput("revenue")
	if err != nil {
		fmt.Println("Revenue cannot be less than or equal to 0.")
		return
	}

	expenses, err := getUserInput("expenses")
	if err != nil {
		fmt.Println("Expenses cannot be less than or equal to 0.")
		return
	}
	taxRate, err := getUserInput("taxRate")
	if err != nil {
		fmt.Println("TaxRate cannot be less than or equal to 0.")
	}

	ebt, eat, ratio := calculateFeatureValues(revenue, expenses, taxRate)

	printValues(ebt, eat, ratio)

	fmt.Println("Saving data to File.")
	saveValuesToFile(ebt, eat, ratio)
}

func getUserInput(variableName string) (float64, error) {
	var userInput float64
	fmt.Printf("Input your %s: \n", variableName)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Amount cannot be negative or 0.")
	}

	return userInput, nil
}

func calculateFeatureValues(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat

	return ebt, eat, ratio
}

func printValues(ebt, eat, ratio float64) {
	fmt.Printf(
		`Total Earnings before tax: %.2f.
		Total Earnings after tax: %.2f.
		Ratio: %.2f.
		`,
		ebt, eat, ratio)
}

func saveValuesToFile(ebt, eat, ratio float64) {
	results := fmt.Sprint("EBT: %.1f \n Ratio: %.1f \n EAT: %.1f", ebt, ratio, eat)
	os.WriteFile(
		accountBalanceFile,
		[]byte(results),
		0644)
}
