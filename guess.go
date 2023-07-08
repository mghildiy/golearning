package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	numGuesses := 5
	fmt.Println("I have chosen a number. Can you guess it?")

	success := false
	for i := 0; i < numGuesses; i++ {
		fmt.Println("You have ", numGuesses-i, "guesses remaining")
		guess := getInputFromUser()
		if guess < target {
			fmt.Println("You guessed ", guess, ", and its LOW")
		} else if guess > target {
			fmt.Println("You guessed ", guess, ", and its HIGH")
		} else {
			success = true
			break
		}
	}

	if success {
		fmt.Println("You won")
	} else {
		fmt.Println("You lost")
	}

}

func getInputFromUser() int {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	num, error := strconv.Atoi(input)
	if error != nil {
		log.Fatal(error)
	}

	return num
}
