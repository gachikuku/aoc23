package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
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
	numberMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	re := regexp.MustCompile(`(\d+)|(zero|one|two|three|four|five|six|seven|eight|nine)`)
	matches := re.FindAllStringSubmatch(line, -1)

	first := 0
	last := 0

	if len(matches) >= 1 {
		if len(matches[0][1]) > 0 {
			firstStr := matches[0][1]
			first = int(firstStr[0] - '0')
		} else if len(matches[0][2]) > 0 {
			first = numberMap[matches[0][2]]
		}
	}

	if len(matches) >= 2 {
		lastMatch := len(matches) - 1
		if len(matches[lastMatch][1]) > 0 {
			lastStr := matches[lastMatch][1]
			last = int(lastStr[len(lastStr)-1] - '0')
		} else if len(matches[lastMatch][2]) > 0 {
			last = numberMap[matches[lastMatch][2]]
		}
	}

	if first == 0 || last == 0 {
		return 0
	} else if first == last {
		return 11 * first
	} else {
		return 10*first + last
	}
}

