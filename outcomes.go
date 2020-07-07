package main

import (
	"fmt"
)

var symptomMap = map[int]string{
	0: "are asymptomatic",
	1: "have moderate illness",
	2: "have serious illness",
}

func symptomEvaluator(n int) int {
	status := 0
	switch {
	case 0 <= n && n <= 59:
		status = 0
	case 60 <= n && n <= 119:
		status = 1
	case n > 120:
		status = 2
	}
	return status
}

func getAge() int {
	fmt.Println("Enter your age: ")
	return readInput()
}

var shameMap = map[int]string{
	1: "GOOD JOB, asshole, you could kill someone. Society shouldn't have to convince you to give" +
		" a shit about other people... 🖕",
	2: "Thanks for wearing a mask but please also keep your distance, eh?",
	3: "Everyone around you thinks you're a considerate, compassionate, superstar team-player! " +
		"Just imagine a world where everyone was this amazing... 🌈",
}

var behaviorMap = map[int]string{
	1: "No mask",
	2: "Mask",
	3: "Mask & physical distancing",
}

var behaviorModifierMap = map[int]int{
	1: 2,
	2: 8,
	3: 10,
}

func getBehavior() int {
	fmt.Printf("Enter a number corresponding to the disease prevention guidelines you follow:\n"+
		"\t(1) %s (2) %s (3) %s:\n", behaviorMap[1], behaviorMap[2], behaviorMap[3])
	num := readInput()
	// TODO: make this recursive?
	if num <= 0 || num > 3 {
		panic(fmt.Sprintf("'%d' is not a valid choice ಠ_ಠ", num))
	}
	fmt.Printf("\nYou chose: %+v\n", behaviorMap[num])
	fmt.Println(shameMap[num] + "\n")
	return num
}
