package main

import "fmt"

func main() {

	revenue := getUserInput("revenue")
	expenses := getUserInput("expenses")
	taxRate := getUserInput("taxRate")

	ebt, eat, ratio := calculateFeatureValues(revenue, expenses, taxRate)

	printValues(ebt, eat, ratio)

}

func getUserInput(variableName string) float64 {
	var userInput float64
	fmt.Printf("Input your %s: \n", variableName)
	fmt.Scan(&userInput)
	return userInput
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
