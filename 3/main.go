package main

import (
	"fmt"
	"gmassman/advent_of_code/utils"
	"log"
	"strconv"
	"strings"
)

type node struct {
	totalClaims int
	claimIDs    []int
}

type grid [][]*node

type elfClaim struct {
	ID, X, Y, W, H int
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var claims []*elfClaim
	maxX, maxY := 0, 0
	for _, line := range lines {
		claim := newElfClaim(line)
		if claim.X+claim.W > maxX {
			maxX = claim.X + claim.W
		}
		if claim.Y+claim.H > maxY {
			maxY = claim.Y + claim.H
		}
		claims = append(claims, claim)
	}

	g := initializeGrid(maxX, maxY)
	// findOverlaps(claims, g)
	findClaimWithoutOverlap(claims, g)
}

func findClaimWithoutOverlap(claims []*elfClaim, g *grid) {
	overlapIDs := make(map[int]bool)
	for _, claim := range claims {
		overlapIDs[claim.ID] = false
		for i := claim.X; i < claim.X+claim.W; i++ {
			for j := claim.Y; j < claim.Y+claim.H; j++ {
				nPtr := (*g)[i][j]
				(*nPtr).totalClaims++
				(*nPtr).claimIDs = append((*nPtr).claimIDs, claim.ID)
			}
		}
	}
	for i := range *g {
		for j := range (*g)[i] {
			nPtr := (*g)[i][j]
			if (*nPtr).totalClaims > 1 {
				for _, id := range (*nPtr).claimIDs {
					overlapIDs[id] = true
				}
			}
		}
	}
	for k, v := range overlapIDs {
		if v == false {
			fmt.Println("no overlap:", k)
		}
	}
}

func findOverlaps(claims []*elfClaim, g *grid) {
	for _, claim := range claims {
		for i := claim.X; i < claim.X+claim.W; i++ {
			for j := claim.Y; j < claim.Y+claim.H; j++ {
				nPtr := (*g)[i][j]
				(*nPtr).totalClaims++
			}
		}
	}
	overlaps := 0
	for i := range *g {
		for j := range (*g)[i] {
			nPtr := (*g)[i][j]
			if (*nPtr).totalClaims > 1 {
				overlaps++
			}
		}
	}
	fmt.Println("total overlaps:", overlaps)
}

func initializeGrid(maxX, maxY int) *grid {
	g := make(grid, maxX+1)
	for i := 0; i <= maxX; i++ {
		g[i] = make([]*node, maxY+1)
		for j := 0; j <= maxY; j++ {
			g[i][j] = &node{0, []int{}}
		}
	}
	return &g
}

func newElfClaim(line string) *elfClaim {
	tokens := strings.Split(strings.TrimLeft(line, "#"), " ")
	id, _ := strconv.Atoi(tokens[0])
	offsets := strings.Split(strings.TrimRight(tokens[2], ":"), ",")
	x, _ := strconv.Atoi(offsets[0])
	y, _ := strconv.Atoi(offsets[1])
	dimensions := strings.Split(tokens[3], "x")
	w, _ := strconv.Atoi(dimensions[0])
	h, _ := strconv.Atoi(dimensions[1])
	return &elfClaim{id, x, y, w, h}
}
