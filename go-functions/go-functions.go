package main

import "fmt"

func main() {
	outputText("Outputing this text message.", "And this is second message to output.")
}

func outputText(str1, str2 string) {
	fmt.Println(str1)
	fmt.Println(str2)
}

// func main() {
// 	const inflation float64 = 3
// 	years := 10.0
// 	expectedReturnRate := 5.5

// 	var investmentAmount float64

// 	fmt.Print("Input amount you want to invest: ")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Print("Input years: ")
// 	fmt.Scan(&years)

// 	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	var futureRealValue = futureValue / math.Pow(1+inflation/100, years)

// 	fmt.Printf("FeatureValue after %d months: %f. \n", years, futureValue)
// 	fmt.Printf("Real FeatureValue after %d months: %f. \n", years, futureRealValue)
// }
