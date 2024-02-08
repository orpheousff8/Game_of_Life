package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	fmt.Printf("%c%c%c%c", upperCase[rand.Intn(len(upperCase))], lowerCase[rand.Intn(len(lowerCase))], number[rand.Intn(len(number))], specialSymbol[rand.Intn(len(specialSymbol))])
}
