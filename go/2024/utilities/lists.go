package utilities

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ExtractLists(path string) ([]int, []int) {

	content, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	listOne := []int{}
	listTwo := []int{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {

		line := scanner.Text()
		locations := strings.Split(line, "   ")

		numOne, err := strconv.ParseInt(strings.TrimSpace(locations[0]), 10, 32)
		if err != nil {
			numOne = 0
		}

		numTwo, err := strconv.ParseInt(strings.TrimSpace(locations[1]), 10, 32)
		if err != nil {
			numTwo = 0
		}

		listOne = append(listOne, int(numOne))
		listTwo = append(listTwo, int(numTwo))

	}

	return listOne, listTwo

}

func Sum(list []int) int {

	total := 0

	for _, number := range list {
		total += number
	}

	return total

}

func NumberOfOccurences(number int, list []int) int {

	total := 0

	for _, element := range list {
		if element == number {
			total += 1
		}
	}

	return total

}
