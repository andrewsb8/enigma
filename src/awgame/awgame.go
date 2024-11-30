package awgame

import (
	"fmt"
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
	Weather     string
	Tiles       []*Tile
}

type Game struct {
	Awmap   Map
	Players []*Player
}

func ParseMapState(game Game) {
	//read through map_state string to find first {
	/*index := 0
	for i := 0; i < len(game.Awmap.Map_state); i++ {
		if game.Awmap.Map_state[i] == '{' {
			index = i
			break
		}
		}*/

	// Split the rest of the string by ;
	// The resulting fields will have the following format
	// [data type]:[size (not for ints)]:data
	// Ex: s:5:"hello" -> string:5:"hello"
	//
	// A line in the map file is then nested
	// AWBWGame
	// - Game Information - ends at buildings below
	// - Player info (CO, etc) - s:7:"players"
	//   - contains entries awbwPlayer
	// - Buildings - s:9:"buildings"
	//   - contains entries of "awbwBuilding"
	// - units - s:5:"units"
	//   - contains entries of "awbwUnit"
	//
	// So, could get substrings for the four categories above and parse individually
	// Or, could parse by ; to get entries as below. Then, split by {, }, and finally :.
	// After all of that splitting, can look for keywords. Maybe break up into functions
	// for map info, players, buildings, and units.
	entries := strings.FieldsFunc(game.Awmap.Map_state, Split)
	for i := 0; i < len(entries); i++ {
		fmt.Printf("%s\n", entries[i])
		vals := strings.Split(entries[i], ":")
		//fmt.Printf("%s %s\n", vals, vals[len(vals)-1])
		if vals[len(vals)-1] == "\"awbwPlayer\"" {
			break
		}
	}
}

func Split(r rune) bool {
	return r == ':' || r == '{' || r == '}' || r == ';'
}
