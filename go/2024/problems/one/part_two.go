package one

import (
	"sort"

	"github.com/mwprogrammer/advent-of-code/go/2024/utilities"
)

func PartTwo() int {

	listOne, listTwo := utilities.ExtractLists("2024/files/day_one.txt")

	sort.Ints(listOne)
	sort.Ints(listTwo)

	similarityScore := []int{}

	for _, number := range listOne {

		similarity := number * utilities.NumberOfOccurences(number, listTwo)
		similarityScore = append(similarityScore, similarity)

	}

	return utilities.Sum(similarityScore)
}
