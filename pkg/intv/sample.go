package main

import (
	"fmt"
	"sort"
)

func main() {
	prevScore := map[string]int{
		"X": 10,
		"Y": 20,
		"Z": 15,
	}

	roundActivity := map[string]activity{
		"X": {"A", 5},
		"Y": {"A", 8},
		"Z": {"C", 3},
	}

	fmt.Printf("%v\n", scoreRound(prevScore, roundActivity, "A"))
}

type activity struct {
	answerChosen string
	timeTaken    int
	// user         string
}

// every user correctly -- 10 points
//
//	user gets 10 , 8 , 6, 4, 2

func scoreRound(prevScore map[string]int, roundActivity map[string]activity, correctAnswer string) map[string]int {
	newScore := make(map[string]int)
	sorted := make(map[int][]string)

	// check correctness
	for user := range prevScore {
		answer := roundActivity[user].answerChosen
		if answer == correctAnswer {
			newScore[user] = prevScore[user] + 10
			if _ , ok := sorted[roundActivity[user].timeTaken]; !ok {
				sorted[roundActivity[user].timeTaken] = []string{}
			}
			sorted[roundActivity[user].timeTaken] = append(sorted[roundActivity[user].timeTaken], user)
		} else {
			newScore[user] = prevScore[user]
		}
	}

	// bucket sort
	alltimes := []int{}
	for k , _ := range sorted {
		alltimes = append(alltimes, k)
	}

	sort.Ints(alltimes)

	bonus := 10
	i := 0
	
	for bonus > 0 && i < len(alltimes) {
		for _, user := range sorted[alltimes[i]] {
			newScore[user] += bonus
			bonus -= 2
			if bonus == 0 {
				return newScore
			}
		}
		i++
	}
	return newScore
}
