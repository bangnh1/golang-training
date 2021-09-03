package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	var inputString string

	for {
		fmt.Println("Enter a number =>0 and <= 100: ")

		fmt.Scan(&inputString)

		guessNum, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if guessNum < 0 || guessNum > 100 {
			fmt.Println("Please enter a number => 0 and <= 100 ")
			continue
		}

		switch {
		case guessNum < x:
			fmt.Println("Your number lower than X ")
		case guessNum > x:
			fmt.Println("Your number greater than X")
		case guessNum == x:
			fmt.Println("You guessed right")
			os.Exit(0)
		}
	}

}
