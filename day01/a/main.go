package main

import (

	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		number := extractNumber(line)
		sum += number
	}

	fmt.Println(sum)
}

func extractNumber(line string) int {

	firstDigit := -1
	lastDigit := -1

	for _, char := range line {
		if char >= '0' && char <= '9' {
			if firstDigit == -1 {
				firstDigit = int(char - '0')
			}
			lastDigit = int(char - '0')
		}
	}

	if firstDigit == -1 || lastDigit == -1 {
		return 0
	} else if firstDigit == lastDigit {
		return 11 * firstDigit
	} else {
        return 10*firstDigit + lastDigit
    }
}

