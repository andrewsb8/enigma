package awgame

import (
	"log"
	"strconv"
	"strings"
)

func ParseMapState(map_state string, game *Game) {
	// Split the map state string by ;, {, and }
	// The resulting fields will have the following format
	// [data type]:[size (not for ints)]:data
	// Ex: s:5:"hello" -> string:5:"hello"
	//
	// The contents of the map state include
	// - awbwGame
	//   - Game Information like weather and funds
	// - Player info (CO, etc) - s:7:"players"
	//   - contains entries awbwPlayer
	// - Buildings - s:9:"buildings"
	//   - contains entries of "awbwBuilding"
	// - units - s:5:"units"
	//   - contains entries of "awbwUnit"
	entries := strings.FieldsFunc(map_state, Split)

	// going to split up entries array into constituent arrays for
	// the categories of map information, player information,
	// unit information, building information. Then can have
	// individual methods for handling each to avoid one long
	// else if loop.
	map_info := SpliceArray(entries, "", "players")
	ParseMapInfo(map_info, game)
	player_info := SpliceArray(entries, "players", "buildings")
	ParsePlayerInfo(player_info, game)
	//building_info := SpliceArray(entries, "buildings", "units")
	//unit_info := SpliceArray(entries, "units", "")
}

func ParseMapInfo(list []string, game *Game) {
	for i := 0; i < len(list); i++ {
		val := ParseString(list[i])
		if val == "weather_type" {
			i += 1 //increment to get data in next line
			game.Weather = ParseString(list[i])
			continue
		} else if val == "day" {
			i += 1
			game.Day = ParseInt(list[i])
			continue
		} else if val == "starting_funds" {
			i += 1
			game.Starting_funds = ParseInt(list[i])
			continue
		}
	}
}

func ParsePlayerInfo(list []string, game *Game) {
	// by design first two entries of list are
	// s:9:"players" and a:[number of players]: .
	// So can just directly parse this info before
	// looping
	game.Num_players = ParseInt(list[1])

	// loop through rest of list
	for i := 2; i < len(list); i++ {
		val := ParseString(list[i])
		if val == "awbwPlayer" {
			game.Players = append(game.Players, &Player{})
			continue
		} else if val == "id" {
			i += 1
			game.Players[len(game.Players)-1].Id = ParseInt(list[i])
			continue
		} else if val == "funds" {
			i += 1
			game.Players[len(game.Players)-1].Funds = ParseInt(list[i])
			continue
		}
	}
	if game.Num_players != len(game.Players) {
		log.Fatal("Number of \"players\" does not match number of \"awbwPlayers\". Please check your input file.")
	}
}

func ParseBuidlingInfo(list []string, game *Game) {
	// need to do two things:
	// - establish which ids are for which captured buildings
	//   - the terrain id is related to the HQ. So if HQ id is 54, base looks like 51.
	// - need to figure out which player to assign the buildings to
}

/*
Take in an array of strings. Each entry has format
[data type]:[size (optional)]:[data]. begin and end
specify the [data] entries to search for. Return subset
of input array starting with entry containing begin and
entry containing end.
*/
func SpliceArray(list []string, begin string, end string) []string {
	begin_index := -1
	end_index := -1
	if begin == "" {
		begin_index = 0
	}
	if end == "" {
		end_index = len(list) - 1
	}

	// only search through string if one or both indices
	// are not determined from above conditionals
	if begin_index == -1 || end_index == -1 {
		// skip first entry which is this: [O 8 "awbwGame" 36 ]
		// for some reason this breaks what I've written
		for i := 1; i < len(list); i++ {
			val := ParseString(list[i])
			if begin_index == -1 && val == begin {
				begin_index = i
			} else if end_index == -1 && val == end {
				end_index = i
				break //can stop iterating if found the end
			}
		}
	}
	return list[begin_index:end_index]
}

/*
Takes in a string of format [data type]:[size]:[data].
Ex of entry: s:5:"hello". Returns string: hello.
*/
func ParseString(entry string) string {
	vals := strings.Split(entry, ":")
	len_vals := len(vals)
	var final string
	if len_vals < 4 {
		// this condition covers data
		// Exs: s:5:"hello", a:2:, i:50
		final = vals[len_vals-1]

	} else {
		// this condition covers the other cases
		// Exs: O:8:"awbwGame":36:,
		// i:0;O:10:"awbwPlayer":30:,
		// i:0;O:8:"awbwUnit":25:,
		// i:33;O:12:"awbwBuilding":8:
		final = vals[len_vals-3]
	}
	if final != "" && final[0] == '"' {
		// remove quotations from string.
		// Ex: "Clear" -> Clear
		final = final[1 : len(final)-1]
	}
	return final
}

/*
Counterpart to ParseString for non-string data
Ex of entry: i:1 or a:1. Both return 1.
*/
func ParseInt(entry string) int {
	vals := strings.Split(entry, ":")
	conversion_index := 1
	// entries starting with "a" look like this before
	// splitting: a:2:, resulting in a 3rd member of the list
	// which is just an empty string, "". Below accounts for
	// this
	if vals[0] == "a" {
		conversion_index = 2
	}
	out, err := strconv.Atoi(vals[len(vals)-conversion_index])
	if err != nil {
		return -10000000
	} else {
		return out
	}
}

func Split(r rune) bool {
	return r == '{' || r == '}' || r == ';'
}
