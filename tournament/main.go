package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type performances struct {
	team        string
	matchPlayed int
	wins        int
	draws       int
	losses      int
	totalPoints int
}

func Tally(reader io.Reader) error {
	teams := make(map[string]*performances)
	r := bufio.NewReader(reader)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ";")

		if len(data) != 3 {
			return errors.New("The content of data does not match what is expected")
		}

		team1, team2, result := data[0], data[1], data[2]

		if _, keyExist := teams[team1]; !keyExist {
			teams[team1] = &performances{team: team1}
		}
		if _, keyExist := teams[team2]; !keyExist {
			teams[team2] = &performances{team: team2}
		}

		switch result {
		case "win":
			teams[team1].matchPlayed++
			teams[team1].wins++
			teams[team1].totalPoints += 3

			teams[team2].matchPlayed++
			teams[team2].losses++
		case "draw":
			teams[team1].matchPlayed++
			teams[team1].draws++
			teams[team1].totalPoints++

			teams[team2].matchPlayed++
			teams[team2].draws++
			teams[team2].totalPoints++

		case "loss":
			teams[team1].matchPlayed++
			teams[team1].losses++

			teams[team2].matchPlayed++
			teams[team2].wins++
			teams[team2].totalPoints += 3
		}

	}
	listOfTeams := []performances{}
	for _, team := range teams {
		listOfTeams = append(listOfTeams, *team)
	}

	sort.Slice(listOfTeams, func(i, j int) bool {
		if listOfTeams[i].totalPoints != listOfTeams[j].totalPoints {
			return listOfTeams[i].totalPoints > listOfTeams[j].totalPoints
		}
		return listOfTeams[i].team < listOfTeams[j].team
	})
	fmt.Printf("%-26s | MP | W | D | L | P \n", "Team")
	for _, team := range listOfTeams {
		fmt.Printf("%-27s| %v  | %v | %v | %v | %v \n", team.team, team.matchPlayed, team.wins, team.draws, team.losses, team.totalPoints)
	}

	return nil
}

func main() {

	data := `Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win`
	Tally(strings.NewReader(data))
}
