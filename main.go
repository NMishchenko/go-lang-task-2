package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const sortingFieldIndex = 0

func main() {
	var content string
	content = ReadFromConsole()

	fmt.Println("Sorted data:\n" + content)
}

func ReadFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	return StartProcessing(scanner)
}

func StartProcessing(scanner *bufio.Scanner) string {
	n := 0
	table := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, ",")

		table = append(table, row)
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][sortingFieldIndex] < table[j][sortingFieldIndex]
	})

	var result strings.Builder

	for _, row := range table {
		result.WriteString(strings.Join(row, ","))
		result.WriteString("\n")
	}

	return result.String()
}