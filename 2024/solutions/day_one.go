package solutions

import (
	"sort"

	"github.com/mwprogrammer/advent-of-code/2024/utilities"
)


func DayOnePartOne() int {

	listOne, listTwo := utilities.ExtractLists("2024/files/day_one.txt")

	sort.Ints(listOne)
	sort.Ints(listTwo)

	distance := []int {}

	for index := range listOne {

		locationOne := listOne[index]
		locationTwo := listTwo[index]

		if locationOne > locationTwo {
			distance = append(distance, locationOne - locationTwo)
		} else {
			distance = append(distance, locationTwo - locationOne)
		}
	}

	return utilities.Sum(distance)
}

func DayOnePartTwo() int {

	listOne, listTwo := utilities.ExtractLists("2024/files/day_one.txt")

	sort.Ints(listOne)
	sort.Ints(listTwo)

	similarityScore := []int {}

	for _, number := range listOne {

		similarity := number * utilities.NumberOfOccurences(number, listTwo)
		similarityScore = append(similarityScore, similarity)

	}
	
	return utilities.Sum(similarityScore)
}