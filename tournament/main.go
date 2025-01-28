package main

import (
	"bufio"
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

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*performances)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		data := strings.Split(line, ";")
		if len(data) != 3 {
			return fmt.Errorf("invalid format: each line must contain exactly 3 fields: %s", line)
		}

		team1, team2, result := data[0], data[1], data[2]

		if result != "win" && result != "loss" && result != "draw" {
			return fmt.Errorf("invalid result '%s': must be 'win', 'loss' or 'draw'", result)
		}

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

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
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

	fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s \n", "Team", "MP", "W", "D", "L", "P")
	for _, team := range listOfTeams {
		fmt.Fprintf(writer, "%-30s| %2d | %2d | %2d | %2d | %2d\n", team.team, team.matchPlayed, team.wins, team.draws, team.losses, team.totalPoints)
	}

	return nil
}

func main() {

	// 	data := `Allegoric Alaskians;Blithering Badgers;win
	// Devastating Donkeys;Courageous Californians;draw
	// Devastating Donkeys;Allegoric Alaskians;win
	// Courageous Californians;Blithering Badgers;loss
	// Blithering Badgers;Devastating Donkeys;loss
	// Allegoric Alaskians;Courageous Californians;win`
	// Tally(strings.NewReader(data))
	fmt.Println(len("Team                             "))
	fmt.Println(len("Devastating Donkeys            "))
	fmt.Println(len("Devastating Donkeys            "))
}

// var errorTestCases = []string{
// 	"Bla;Bla;Bla",
// 	"Devastating Donkeys_Courageous Californians;draw",
// 	"Devastating Donkeys@Courageous Californians;draw",
// 	"Devastating Donkeys;Allegoric Alaskians;dra",
// }
