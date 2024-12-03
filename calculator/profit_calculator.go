// package main

// import "fmt"

// func main() {

// 	revenue, expenses, taxRate := 0.0, 0.0, 0.0

// 	fmt.Println("Input your revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Println("Input your expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Println("Input your taxRate: ")
// 	fmt.Scan(&taxRate)

// 	ebt := revenue - expenses
// 	var eat = ebt * (1 - taxRate/100)

// 	var ratio = ebt / eat

// 	// formatedEbt := fmt.Sprintf("Total Earnings before tax: %.2f. \n", ebt)
// 	// formatedEat := fmt.Sprintf("Total Earnings after tax: %.2f. \n", eat)
// 	// formatedRatio := fmt.Sprintf("Ratio: %.2f. \n", ratio)

// 	// fmt.Print(formatedEbt, formatedEat, formatedRatio)

// 	fmt.Printf(
// 		`Total Earnings before tax: %.2f.
// 		Total Earnings after tax: %.2f.
// 		Ratio: %.2f.
// 		`,
// 		ebt, eat, ratio)
// }
