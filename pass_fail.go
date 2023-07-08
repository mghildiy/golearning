package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	fmt.Println(input)
	grade, error := strconv.ParseFloat(input, 64)
	if error != nil {
		log.Fatal(error)
	}

	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println(status)

}
