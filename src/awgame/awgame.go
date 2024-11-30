package awgame

import (
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	Funds                 int
	Units                 []*Unit
	Army_value            int
	Income                int
	Num_income_properties int
	Num_properties        int //including labs and comm towers
}

type Unit struct {
	Unit_id      int
	Unit_type    string
	Movement     int
	Vision       int
	X_position   int
	Y_position   int
	Ammo         int
	Gas          int
	Can_capture  bool
	Is_capturing bool
}

type Tile struct {
	Terrain_id    int
	Terrain_type  string
	Can_capture   bool
	Capture_value int
	X_location    int
	Y_location    int
}

type Map struct {
	Map_state   string
	Map_width   int
	Map_height  int
	Num_players int
	Has_hq      bool
	Tiles       []*Tile
}

type Game struct {
	Awmap          Map
	Num_players    int
	Players        []*Player
	Day            int
	Fog            bool
	Funds          int
	Starting_funds int
	Weather        string
}

func ParseMapState(game Game) {
	// Split the string by ;, {, and }
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
	entries := strings.FieldsFunc(game.Awmap.Map_state, Split)
	for i := 0; i < len(entries); i++ {
		vals := strings.Split(entries[i], ":")
		if vals[len(vals)-1] == "\"weather_type\"" {
			i += 1 //increment to get data in next line
			game.Weather = ParseString(entries[i])
		} else if vals[len(vals)-1] == "\"day\"" {
			i += 1
			game.Day = ParseInt(entries[i])
		} else if vals[len(vals)-1] == "\"players\"" {
			i += 1
			game.Num_players = ParseInt(entries[i])
		} else if vals[len(vals)-1] == "\"starting_funds\"" {
			i += 1
			game.Starting_funds = ParseInt(entries[i])
		}
	}
	fmt.Printf("%d %d %d %s\n", game.Num_players, game.Starting_funds, game.Day, game.Weather)
}

/*
Takes in a string of format [data type]:[size]:[data].
Ex of entry: s:5:"hello". Returns string: hello.
*/
func ParseString(entry string) string {
	vals := strings.Split(entry, ":")
	final := string(vals[len(vals)-1])
	return final[1 : len(final)-1] //remove quotations from string
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
