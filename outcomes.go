package main

import (
	"fmt"
	"os"
	"time"
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
	case n >= 120:
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

func recoveryEvaluator(n int) bool {
	var status bool
	if 0 <= n && n <= 119 {
		status = true
	}
	return status
}

func recovery(hp, originalHP int, timeframe string) {
	fmt.Println("You're on the road to recovery!\n")
	fmt.Printf("You slowly, painfully regain 1 health per %s, possibly taking up to 6 %ss...\n", timeframe,
		timeframe)
	var count int
	for i := 1; i <= 6; i++ {
		if hp == originalHP {
			break
		}
		if timeframe == "week" {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(2 * time.Second)
		}
		fmt.Printf("\t%d\n", i)
		hp++
		count++
	}
	fmt.Printf("...After %d %s(s) you've healed back up to %d.\n", count, timeframe, hp)
	if hp > originalHP {
		hp = originalHP
	}
	if hp < originalHP {
		fmt.Println("\nIs that less than what you started with? Yeah, this disease is terrible." +
			" Many people just like you have been left with long-term, permanent organ damage." +
			" But, hey, at least you're still alive!\n")
	} else {
		fmt.Println("\nYou've managed to recover fully with no long-lasting effects. Count " +
			"yourself incredibly lucky!")
	}
	os.Exit(0)
}

func hospitalEvaluator(n int) bool {
	var result bool
	if 0 <= n && n <= 109 {
		result = true
	}
	return result
}

func goToICU(age int) int {
	fmt.Println("Uh-oh, you're in the ICU. Things could get worse...\n")
	result := rollDice(1, 100) + age
	fmt.Printf("You rolled (d100 + your age) = %d\n", result)
	return result
}

func goOnVentilator(age int) int {
	fmt.Println("The hospital staff is putting you on a ventilator. Hopefully you got to say " +
		"goodbye to your loved ones...\n")
	result := rollDice(1, 100) + age
	fmt.Printf("You rolled (d100 + your age) = %d\n", result)
	return result
}
