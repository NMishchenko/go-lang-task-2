package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"strconv"
)

const sortingFieldIndex = 0

func main() {
	var content string
	content = ReadFromConsole()

	if content != "" {
		fmt.Println("Sorted data:\n" + content)

		dateTimeNow := time.Now()
		fileName := strconv.Itoa(dateTimeNow.Year()) + "-" + dateTimeNow.Month().String() + "-" + strconv.Itoa(dateTimeNow.Day()) + "_" + strconv.Itoa(dateTimeNow.Hour()) + "-" + strconv.Itoa(dateTimeNow.Minute()) + "-" + strconv.Itoa(dateTimeNow.Second()) + ".csv"
		WriteToFile(content, fileName)
	}
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

		if n == 0 {
			n = len(row)
		}

		if line == "" {
			break
		}

		if n != len(row) {
			log.Fatalf("Row has %d columns, but must have %d\n", len(row), n)
		}

		table = append(table, row)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
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

func WriteToFile(content, fileName string) {
	if fileName != "" {
		file, err := os.Create(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		_, err = file.WriteString(content)

		if err != nil {
			log.Fatal(err)
		}
	}
}