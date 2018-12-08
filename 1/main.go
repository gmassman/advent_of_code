package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// finalFreq(scanner)

	numbers, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	seen := make(map[int]bool)
	start := 0
	for {
		start = sumWithDupCheck(&seen, start, numbers)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sumWithDupCheck(seen *map[int]bool, freq int, numbers []string) int {
	for _, str := range numbers {
		number, _ := strconv.Atoi(str[1:])
		switch str[0] {
		case '+':
			freq += number
		case '-':
			freq -= number
		default:
			log.Fatal("can't parse the line")
		}
		if (*seen)[freq] == true {
			fmt.Printf("Found first duplicate freqency: %v\n", freq)
			log.Fatalf("program over with duplicate %v\n", freq)
		}
		(*seen)[freq] = true
	}

	return freq
}

// sum the frequency values from the scanner into freq
func finalFreq(freq int, s *bufio.Scanner) int {
	for s.Scan() {
		str := s.Text()
		number, _ := strconv.Atoi(str[1:])
		switch str[0] {
		case '+':
			freq += number
		case '-':
			freq -= number
		default:
			log.Fatal("can't parse the line")
		}
		fmt.Println(freq)
	}
	fmt.Printf("final frequency: %v\n", freq)
	return freq
}
