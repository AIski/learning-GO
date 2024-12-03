package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation float64 = 3
	years := 10.0
	expectedReturnRate := 5.5

	var investmentAmount float64

	fmt.Print("Input amount you want to invest: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Input years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFeatureValues(investmentAmount, expectedReturnRate, years, inflation)

	fmt.Printf("FeatureValue after %d months: %f. \n", years, futureValue)
	fmt.Printf("Real FeatureValue after %d months: %f. \n", years, futureRealValue)
}

func calculateFeatureValues(
	investmentAmount float64,
	expectedReturnRate float64,
	years float64,
	inflation float64) (float64, float64) {

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	return futureValue, futureRealValue
}

// THIS IS THE SHORTERWAY, YOU CAN DECLARE VARIABLE IN RETURN DECLARATION
// func calculateFeatureValues(
// 	investmentAmount float64,
// 	expectedReturnRate float64,
// 	years float64,
// 	inflation float64) (fv float64, rfv float64) {

// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflation/100, years)
// 	return fv,rfv
// }
