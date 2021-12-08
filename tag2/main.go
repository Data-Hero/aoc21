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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	task(lines)
	task2(lines)

}

func task2(lines []string) {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		
		direction, step := lineSplit[0], lineSplit[1]
		
		switch direction {
		case "forward":
			value, _ := strconv.Atoi(step)
			horizontalPosition += value
			depth += value * aim
		case "up":
			value, _ := strconv.Atoi(step)
			aim -= value
		case "down":
			value, _ := strconv.Atoi(step)
			aim += value
		}
	}
	fmt.Println(depth*horizontalPosition)

}

func task(lines []string) {
	horizontalPosition := 0
	depth := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		
		direction, step := lineSplit[0], lineSplit[1]
		
		switch direction {
		case "forward":
			value, _ := strconv.Atoi(step)
			horizontalPosition += value
		case "up":
			value, _ := strconv.Atoi(step)
			depth -= value
		case "down":
			value, _ := strconv.Atoi(step)
			depth += value
		}
	}
	fmt.Println(depth*horizontalPosition)
}
