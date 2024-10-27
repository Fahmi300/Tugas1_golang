package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input String: ")
	sent, _ := reader.ReadString('\n') 

	words := strings.Split(sent, " ")

	fmt.Println(words)
	var reversedWords string
	var output []string
	for i := 0; i <= len(words) - 1; i++ {	
		for j := len(words[i]) - 1; j >= 0; j-- {
			reversedWords += string(words[i][j])
		}
		output = append(output, reversedWords)
		reversedWords = ""
	}

	fmt.Println("Output String:", strings.Join(output, " "))
}
