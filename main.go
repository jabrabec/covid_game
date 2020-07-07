package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(fmt.Sprintf("'%s' is not a valid number ಠ_ಠ", text))
	}
	return num
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\t *** LET'S PLAY THE COVID GAME! ***")

	age := getAge()
	originalHP := 20
	hp := 20
	fmt.Printf("You are %d years old and currently have %d health. Great! Let's continue...\n\n",
		age, hp)

	fmt.Printf("Uh-oh, you caught COVID-19. How does it impact you?\n")

	symptomCheck := rollDice(1, 100) + age
	fmt.Printf("You rolled (d100 + your age) = %d\n", symptomCheck)
	symptoms := symptomEvaluator(symptomCheck)
	fmt.Printf("You %s!\n\n", symptomMap[symptoms])

	switch symptoms {
	case 0:
		fmt.Println("You can still carry the disease and infect other people.")
		fmt.Println("What precautions to you take to protect others?")
		behavior := getBehavior()
		infected := rollDice(2, 6) - behaviorModifierMap[behavior]
		if infected < 0 {
			infected = 0
		}
		fmt.Printf("You rolled (2d6 - %d) = %d\n", behaviorModifierMap[behavior], infected)
		fmt.Printf("Your choices lead you to infect %d other people. Can you live with that?\n",
			infected)
	case 1:
		lostHP := rollDice(2, 4)
		fmt.Printf("You rolled 2d4 = %d\n", lostHP)
		hp -= lostHP
		fmt.Printf("Your moderate case causes you to lose %d health. You now have %d.\n", lostHP,
			hp)
		fmt.Printf("You slowly, painfully regain 1 health per week for the next 6 weeks...\n")
		for i := 1; i <= 6; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("\t%d\n", i)
		}
		hp += 6
		if hp > originalHP {
			hp = originalHP
		}
		fmt.Printf("After 6 weeks you've healed back up to %d.\n", hp)
		if hp < originalHP {
			fmt.Println("Is that less than what you started with? Yeah, this disease is terrible." +
				" Many people just like you have been left with long-term, permanent organ damage." +
				" But, hey, at least you're still alive!\n")
		} else {
			fmt.Println("You've managed to recover fully with no long-lasting effects. Count " +
				"yourself incredibly lucky!")
		}

	case 2:
	}
}
