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

		if line == "" {
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
	caloryLength := len(foodCalories)

	fmt.Println()
	fmt.Println(getMaxCalories(foodCalories, caloryLength))
	sortCalories(foodCalories, caloryLength)
	fmt.Println("top 3 : ")
	var sum = 0
	for i := caloryLength - 1; i >= caloryLength-3; i-- {
		sum += foodCalories[i]
		fmt.Println(foodCalories[i])
	}
	fmt.Println(sum)

}

func getMaxCalories(foodCalories []int, length int) int {
	max := 0
	for i := 0; i < length; i++ {
		if max < foodCalories[i] {
			max = foodCalories[i]
		}
	}
	return max
}

func sortCalories(calories []int, length int) {
	swap := 0
	for i := 0; i < length; i++ {
		swapped := false

		for j := 0; j < length-i-1; j++ {
			if calories[j] > calories[j+1] {
				swap = calories[j]
				calories[j] = calories[j+1]
				calories[j+1] = swap
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}
}
