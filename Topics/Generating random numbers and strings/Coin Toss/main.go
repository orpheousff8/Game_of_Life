package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	_, err := fmt.Scanln(&seed)
	if err != nil {
		return
	}
	rand.Seed(seed)

	if rand.Float64() < 0.5 {
		fmt.Println("TAILS")
	} else {
		fmt.Println("HEADS")
	}

}
