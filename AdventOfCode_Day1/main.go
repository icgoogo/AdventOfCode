package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal("error when opening the file", err)
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal("erro when scanning file", err)
	}

	var foodCalories []int
	elfCalorySum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if isEmptyLine(line) {
			foodCalories = append(foodCalories, elfCalorySum)
			elfCalorySum = 0
		} else {
			calory, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("conversion error", err)
			}
			elfCalorySum += calory
		}
	}

	if elfCalorySum != 0 {
		foodCalories = append(foodCalories, elfCalorySum)
	}

	max := 0
	for i := 0; i < len(foodCalories); i++ {
		fmt.Println(foodCalories[i])
		if max < foodCalories[i] {
			max = foodCalories[i]
		}
	}
	fmt.Println()
	fmt.Print(max)
}

func isEmptyLine(line string) bool {
	return line == ""
}
