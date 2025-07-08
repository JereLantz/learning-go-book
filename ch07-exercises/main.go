package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

/* 1.
You have been asked to manage a basketball league and are going to write a program to help you.
Define two types.
The first one, called Team, has a field for the name of the team and a field for the player names.
The second type is called League and has a field called Teams for the teams in
the league and a field called Wins that maps a team’s name to its number of wins.
*/
type Team struct {
	Name        string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

/* 2.
Add two methods to League.
The first method is called MatchResult.
It takes four parameters: the name of the first team, its score in the game,
the name of the second team, and its score in the game.
This method should update the Wins field in League.
Add a second method to League called Ranking that returns a slice of the team names in order of wins.
Build your data structures and call these methods from the main function in your program using some sample data.
*/
func (l *League)MatchResult (team1Name string, team1Score int, team2Name string, team2Score int){
	if _, ok := l.Teams[team1Name]; !ok{
		return
	}
	if _, ok := l.Teams[team2Name]; !ok{
		return
	}

	if team1Score < team2Score {
		l.Wins[team2Name]++
	} else if team1Score > team2Score {
		l.Wins[team1Name]++
	}
}

func (l League)Ranking () []string{
	names := make([]string, 0, len(l.Teams))
	for team := range l.Teams{
		names = append(names, team)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

/* 3.
Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
Call this function from main.
*/

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, row := range r.Ranking() {
		io.WriteString(w, row)
		w.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)

	RankPrinter(l, os.Stdout)
}
