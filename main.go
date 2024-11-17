package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var num float32 = 523525.5525125
	fmt.Println(num)

	var str string = "To jest zwykły string"
	fmt.Println(str)

	fmt.Print("Długość stringa w runach):")

	fmt.Print(utf8.RuneCountInString(str))

	var myRune rune = 'a'

	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)
}
