package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
    "strconv"
)

func gameID(conudrum string) int {
    gameRegex := regexp.MustCompile(`Game (\d+)`)
    gameMatch := gameRegex.FindStringSubmatch(conudrum)
    gameNumber := gameMatch[1]
    result, err := strconv.Atoi(gameNumber)
    if err != nil {
        panic(err)
    } 
    return result
}

func possible(conudrum string) bool {
    colorRegex := regexp.MustCompile(`(\d+) (\w+)`)
    colorMatches := colorRegex.FindAllStringSubmatch(conudrum, -1)
    for _, match := range colorMatches {
        colorNumber := match[1]
        number, err := strconv.Atoi(colorNumber)
        if err != nil {
            panic(err)
        }
        colorName := match[2]
        // here is the possible culprit
        switch colorName {
        case "red":
            if number > 12 { return false }
        case "green":
            if number > 13 { return false }
        case "blue":
            if number > 14 { return false }
        }
    }
    return true 
}

func main() {

    var ans int

    readFile, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        if possible(fileScanner.Text()) == true {
            ans += gameID(fileScanner.Text())
        }
    }

    readFile.Close()

    fmt.Println(ans)
}


