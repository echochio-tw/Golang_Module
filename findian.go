package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter string:=>")
	inputReader := bufio.NewReader(os.Stdin)
	sourceString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	sourceString = sourceString[:len(sourceString)-2]
	sourceString = strings.ToLower(sourceString)
	if strings.HasPrefix(sourceString, "i") &&
		strings.HasSuffix(sourceString, "n") &&
		strings.Contains(sourceString, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
