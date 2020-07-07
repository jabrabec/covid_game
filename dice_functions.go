package main

import "math/rand"

func rollDice(numRolls, dieSides int) int {
	var result int
	for i := 0; i < numRolls; i++ {
		result += rand.Intn(dieSides + 1)
	}
	return result
}
