package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type Player struct {
	name string
	age  int
}

type Team struct {
	name    string
	players []Player
}

type Wins map[string]int

type League struct {
	Teams []Team
	Wins
}

func (l *League) MatchResult(
	ftName string, ftMatchScore int, stName string, stMatchScore int) {

	if ftMatchScore > stMatchScore {
		l.Wins[ftName]++
	} else if ftMatchScore < stMatchScore {
		l.Wins[stName]++
	} else {
		return
	}
}

func (l *League) Ranking() []string {
	result := make([]string, len(l.Teams))
	for i, team := range l.Teams {
		result[i] = team.name
	}

	sort.Slice(result, func(i, j int) bool {
		return l.Wins[result[i]] > l.Wins[result[j]]
	})

	return result
}

func main() {
	team1 := Team{"manu", []Player{}}
	team2 := Team{"asn", []Player{}}
	premierLeague := League{[]Team{team1, team2}, Wins{}}

	for i := 0; i < 6; i++ {
		premierLeague.
			MatchResult(
				team1.name, getRandomScore(5), team2.name, getRandomScore(5))
	}
	fmt.Println(premierLeague.Wins, ": ", premierLeague.Ranking())
}

func getRandomScore(rng int) int {
	return r.Intn(rng)
}
