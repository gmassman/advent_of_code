package main

import (
	"fmt"
	"gmassman/advent_of_code/utils"
	"log"
)

func main() {
	// boxIDs, err := readLines("input.txt")
	boxIDs, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// result := calculateChecksum(boxIDs)
	// fmt.Printf("final checksum: %v\n", result)
	countMap := make(map[string]*letterCounts)
	for _, a := range boxIDs {
		countMap[a] = newLetterCounts(a)
	}

	distances := make(map[int][]string)
	for i := 0; i < len(boxIDs); i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			d := letterCountsDistance(countMap[boxIDs[i]], countMap[boxIDs[j]])
			if _, ok := distances[d]; ok {
				distances[d] = append(distances[d], fmt.Sprintf("%s:%s", boxIDs[i], boxIDs[j]))
			} else {
				distances[d] = []string{fmt.Sprintf("%s:%s", boxIDs[i], boxIDs[j])}
			}
		}
	}
	for k, v := range distances {
		fmt.Println(k, len(v))
	}
	fmt.Println(distances[1])
}

type letterCounts map[rune]int

func letterCountsDistance(a, b *letterCounts) int {
	diff := 0
	for ka, va := range *a {
		if vb, ok := (*b)[ka]; ok {
			diff += utils.Max((va - vb), 0)
		} else {
			diff += va
		}
	}
	return diff
}

func newLetterCounts(id string) *letterCounts {
	letters := make(letterCounts)
	for _, c := range id {
		if _, ok := letters[c]; !ok {
			letters[c] = 1
		} else {
			letters[c]++
		}
	}

	return &letters
}

func calculateChecksum(boxIDs []string) int {
	twoCounts, threeCounts := 0, 0
	for _, id := range boxIDs {
		chars := make(letterCounts)
		for _, c := range id {
			if val, ok := chars[c]; !ok {
				chars[c] = 1
			} else {
				chars[c] = val + 1
			}
		}

		for _, v := range chars {
			if v == 2 {
				twoCounts++
				break
			}
		}
		for _, v := range chars {
			if v == 3 {
				threeCounts++
				break
			}
		}
	}
	return twoCounts * threeCounts
}
