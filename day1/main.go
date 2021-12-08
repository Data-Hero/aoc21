package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

}

func task(lines []string) {
	var count int
	var last int
	for i, line := range lines {
		if i == 0 {
			last, _ = strconv.Atoi(line)
			continue
		} else {
			lineint, _ := strconv.Atoi(line)
			if last < lineint {
				count++
			}
			last, _ = strconv.Atoi(line)
		}
	}
	fmt.Println(count)
}
