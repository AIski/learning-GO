package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation float64 = 3
	investmentAmount, years := 1000.0, 10.0
	expectedReturnRate := 5.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflation/100, years)

	fmt.Printf("FeatureValue after %d months: %f. \n", years, futureValue)
	fmt.Printf("Real FeatureValue after %d months: %f. \n", years, futureRealValue)
}
