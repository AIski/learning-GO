//go mod init <app-name>  // creates new go module, base structure of an app.structure
//go run main.go // runs the app

all code must belong to a package - first line , like : 
//package main

application for it to start needs an entry point - main() method:
// func main() {
// }


for printing code to console, we need to import fmt package
import (
	"fmt"
    }
    then we can just use fmt.Print("Message");

// const is for constants, they are like variables, just their value cannot be changed.