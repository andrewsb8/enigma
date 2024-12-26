package awgame

import (
	"log"
	"strconv"
	"strings"
)

func ParseTerrain(terrain []string, game *Game) {
	// terrain file ends with empty line, so -1
	game.Awmap.Map_height = len(terrain) - 1
	for i := 0; i < game.Awmap.Map_height; i++ {
		game.Awmap.Tiles = append(game.Awmap.Tiles, []*Tile{})
		row := strings.Split(terrain[i], ",")
		if i == 0 {
			game.Awmap.Map_width = len(row)
		}
		for j := 0; j < len(row); j++ {
			game.Awmap.Tiles[i] = append(game.Awmap.Tiles[i], CreateTile(row[j]))
		}
	}
}

func CreateTile(terrain_id string) *Tile {
	terrain_map := getTerrainTypeMap()
	tile := *GetTerrainFromID(terrain_id, &Tile{}, terrain_map)
	return &tile
}

/*
Parse neutral terrain from map file. All of the nations have unique
codes for their properties once captured. These IDs can be parsed from
the map state file.
*/
func GetTerrainFromID(terrain_id string, tile *Tile, terrain_map map[int]*Tile) *Tile {
	int_id, err := strconv.Atoi(terrain_id)
	if err != nil {
		log.Fatal("Invalid terrain id. Check terrain file.")
	}

	if int_id >= 4 && int_id <= 14 {
		// rivers
		tile = terrain_map[4]
	} else if int_id >= 15 && int_id <= 27 {
		// roads
		tile = terrain_map[15]
	} else if int_id >= 29 && int_id <= 32 {
		// shoals
		tile = terrain_map[29]
	} else if (int_id >= 101 && int_id <= 110) || int_id == 113 || int_id == 114 {
		// pipes
		tile = terrain_map[101]
	} else if int_id == 111 || int_id == 112 {
		// silos
		tile = terrain_map[111]
	} else if (int_id >= 38 && int_id <= 100) || (int_id >= 117 && int_id <= 194) || (int_id >= 196 && int_id <= 216) {
		// all captured property ranges
		// will use different map here
		tile = &Tile{}
	} else {
		// all terrain which only has one id
		// plain, mountain, sea, reefs
		tile = terrain_map[int_id]
	}
	tile.Terrain_id = int_id
	return tile
}

/*
returns a map of Tile types for regular terrain and
neutral properties. int is terrain id
*/
func getTerrainTypeMap() map[int]*Tile {
	terrainMap := map[int]*Tile{
		1: {
			Terrain_type:        "plain",
			Defense_stars:       1,
			Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		2: {
			Terrain_type:        "mountain",
			Defense_stars:       4,
			Movement_cost_clear: [8]int{2, 1, 100, 100, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		3: {
			Terrain_type:        "woods",
			Defense_stars:       2,
			Movement_cost_clear: [8]int{1, 1, 2, 3, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		4: {
			Terrain_type:        "river",
			Defense_stars:       0,
			Movement_cost_clear: [8]int{2, 1, 100, 100, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		15: {
			Terrain_type:        "road",
			Defense_stars:       0,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		28: {
			Terrain_type:        "sea",
			Defense_stars:       1,
			Movement_cost_clear: [8]int{100, 100, 100, 100, 1, 1, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		29: {
			Terrain_type:        "shoals",
			Defense_stars:       0,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 1, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		33: {
			Terrain_type:        "reefs",
			Defense_stars:       1,
			Movement_cost_clear: [8]int{100, 100, 100, 100, 2, 2, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		101: {
			Terrain_type:        "pipe",
			Defense_stars:       0,
			Movement_cost_clear: [8]int{100, 100, 100, 100, 100, 100, 100, 1},
			Can_capture:         false,
			Capture_points:      -1,
		},
		111: {
			Terrain_type:        "silo",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         false,
			Capture_points:      -1,
		},
		34: {
			Terrain_type:        "city",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
		35: {
			Terrain_type:        "base",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
		36: {
			Terrain_type:        "airport",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
		37: {
			Terrain_type:        "port",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
		133: {
			Terrain_type:        "comm tower",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
		145: {
			Terrain_type:        "lab",
			Defense_stars:       3,
			Movement_cost_clear: [8]int{1, 1, 1, 1, 100, 100, 1, 100},
			Can_capture:         true,
			Capture_points:      20,
		},
	}
	return terrainMap
}
