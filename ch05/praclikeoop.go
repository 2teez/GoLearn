package main

type wins int
type score int

type Team struct {
	TeamName string
	Players  []string
}

type League struct {
	Teams []Team
	Wins  map[string]wins
}

func (l *League) MatchResult(team1 Team, team1GameScore score,
	team2 Team, team2GameScore score) {
	l.Teams = append(l.Teams, team1, team2)
	l.Wins[team1.TeamName] = wins(team1GameScore)
	l.Wins[team2.TeamName] = wins(team2GameScore)
}

func (l *League) Ranking() []string {
	result := []string{}
	//for _, name := range l.Teams {

	//}
	return result
}
