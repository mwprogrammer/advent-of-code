package file

import (
	"bufio"
	"os"
	"strings"
)

// ReadRows() allows for one to read an n number of rows in a file.
// If count is 0, all rows available in a file are returned.
func ReadRows(path string, count int) ([]string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	rows := make([]string, count)

	scanner := bufio.NewScanner(file)

	for i := 0; i < count; i++ {
		rows = append(rows, scanner.Text())
	}

	return rows, nil

}

// ReadColumns() allows for one to read an n number of columns in a file.
// If count is 0, all columns available in a file are returned.
func ReadColumns(path string, count int, separator string) ([][]string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	columns := make([][]string, count)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		row := strings.Split(scanner.Text(), separator)

		for i := 0; i < count; i++ {

			column := columns[i]
			column = append(column, row[i])

			columns[i] = column

		}
	}

	return columns, nil

}
