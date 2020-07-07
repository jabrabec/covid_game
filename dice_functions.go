package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func rollDice(numRolls, dieSides int) int {
	fmt.Printf("Press 'Enter' to roll %dd%d:", numRolls, dieSides)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	var result int
	for i := 0; i < numRolls; i++ {
		result += rand.Intn(dieSides + 1)
	}
	return result
}
