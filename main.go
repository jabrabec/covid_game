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

	fmt.Println("\n\t *** LET'S PLAY THE COVID GAME! ***\n")

	age := getAge()
	originalHP := 20
	hp := 20
	fmt.Printf("You are %d years old and currently have %d health. Great! Let's continue...\n\n",
		age, hp)

	fmt.Println("Uh-oh, you caught COVID-19. How does it impact you?\n")

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
		fmt.Printf("You rolled (2d6 - %d) = %d\n", behaviorModifierMap[behavior], infected)
		if infected < 0 {
			infected = 0
		}
		fmt.Printf("Your choices lead you to infect %d other people.", infected)
		if infected == 0 {
			fmt.Printf(" WELL DONE!")
		} else {
			fmt.Printf(" Can you live with that?\n")
		}

	case 1:
		lostHP := rollDice(2, 4)
		fmt.Printf("You rolled 2d4 = %d\n", lostHP)
		hp -= lostHP
		fmt.Printf("Your moderate case causes you to lose %d health. You now have %d.\n", lostHP,
			hp)
		recovery(hp, originalHP, "week")

	case 2:
		lostHP := rollDice(2, 6)
		fmt.Printf("You rolled 2d6 = %d\n", lostHP)
		hp -= lostHP
		fmt.Printf("Your serious case causes you to lose %d health. You now have %d.\n\n", lostHP,
			hp)

		recoveryCheck := rollDice(1, 100) + age
		fmt.Printf("You rolled (d100 + your age) = %d\n", recoveryCheck)

		if recoveryEvaluator(recoveryCheck) {
			recovery(hp, originalHP, "month")
		}

		time.Sleep(1 * time.Second)

		icuResult := goToICU(age)
		if hospitalEvaluator(icuResult) {
			fmt.Println("You're fortunate enough to make it out of the ICU. Do you have good " +
				"(any?) health insurance? Do you live in an area where medical resources aren't " +
				"currently being rationed? Many others are not so lucky...\n")
			recovery(hp, originalHP, "month")
		}

		time.Sleep(1 * time.Second)

		ventilatorResult := goOnVentilator(age)
		if hospitalEvaluator(ventilatorResult) {
			fmt.Println("You're fortunate enough to be taken off of the ventilator. Things are " +
				"finally looking up for you!\n")
			recovery(hp, originalHP, "month")
		} else {
			fmt.Println("\nWELLP, you have died of COVID-19. But at least those people were able" +
				" to get their hair cut, right? ¯\\_(ツ)_/¯")
		}
	}
}
