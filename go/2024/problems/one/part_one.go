package one

import (
	"log/slog"
	"strconv"

	"github.com/mwprogrammer/advent-of-code/go/2024/utilities/file"
)

// Part one: Get the distance between each two locations and then return the sum of all distances.
func PartOne() int {

	locations, err := file.ReadColumns("input.txt", 2, " ")

	if err != nil {
		slog.Error(err.Error())
		return 0
	}

	total_distance := 0

	for i := 0; i < len(locations[0]); i++ {

		locationOne, err := strconv.Atoi(locations[0][i])

		if err != nil {
			continue
		}

		locationTwo, err := strconv.Atoi(locations[1][i])

		if err != nil {
			continue
		}

		total_distance += locationOne + locationTwo

	}

	return total_distance
}
