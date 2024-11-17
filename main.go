package main

import (
	"fmt"
)

func main() {

	const busTicketsTotal = 10

	fmt.Printf("Total tickets bus limit we got: %v.\n", busTicketsTotal)

	var currentTickets = busTicketsTotal
	currentTickets = currentTickets - 1
	currentTickets = currentTickets - 1

	fmt.Printf("Current tickets we got: %v from total of %v tickets.\n", currentTickets, busTicketsTotal)

}
