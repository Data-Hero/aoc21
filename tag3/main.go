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
	task2(lines)
}

func task2(lines []string) {
	oxy := lines
	co2 := lines
	
	i := 0
	for {
		gamma, _ := count(oxy)
		oxy = filter(oxy, i, ([]rune(gamma))[i])
		if len(oxy) < 2 {
			break
		}
		i++
	}


	i = 0
	for {
		_, epsilon := count(co2)
		co2 = filter(co2, i, ([]rune(epsilon))[i])
		fmt.Println(i, co2)

		if len(co2) < 2 {
			break
		}
		i++
	}

	a, _ := strconv.ParseInt(oxy[0], 2, 64)
	b, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println(a * b)
}

func task(lines []string) {
	gamma, epsilon := count(lines)

	gamma10, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon10, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gamma10 * epsilon10)
}

func count(lines []string) (string, string) {
	length := len([]rune(lines[0]))
	zeroes := make([]int, length)
	ones := make([]int, length)

	for _, line := range lines {
		for j, letter := range line {
			if string(letter) == "1" {
				ones[j]++
			} else {
				zeroes[j]++
			}
		}
	}

	var gamma string
	var epsilon string

	for i, _ := range zeroes {
		if zeroes[i] <= ones[i] {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	return gamma, epsilon
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func filter(vals []string, position int, filterBy rune) []string {
	var result []string
	for _, val := range vals {
		if ([]rune(val))[position] == filterBy {
			result = append(result, val)
		}
	}

	return result
}
