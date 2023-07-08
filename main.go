// main.go file
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello Go")
	x := time.Now()
	fmt.Println(x)
	fmt.Println(x.Day())

	str1 := "Hello ?"
	replacer := strings.NewReplacer("?", "World")
	fmt.Println(replacer.Replace(str1))
	fmt.Println(str1)

}
