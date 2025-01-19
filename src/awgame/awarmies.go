package awgame

/*
returns a map of Tile types for captured properties
the first int in map will be country code and second
one will be terrain id.
*/
func GetArmyPropertyMap() map[int]map[int]*Tile {
	armyMap := map[int]map[int]*Tile{
		4: { //yellow comet
			53: {
				Terrain_type:        "city",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			54: {
				Terrain_type:        "base",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			55: {
				Terrain_type:        "airport",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			56: {
				Terrain_type:        "port",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			57: {
				Terrain_type:        "hq",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			136: {
				Terrain_type:        "comm tower",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			148: {
				Terrain_type:        "lab",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
		},
		20: { //purple lightning
			172: {
				Terrain_type:        "city",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			171: {
				Terrain_type:        "base",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			170: {
				Terrain_type:        "airport",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			176: {
				Terrain_type:        "port",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			174: {
				Terrain_type:        "hq",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			173: {
				Terrain_type:        "comm tower",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
			175: {
				Terrain_type:        "lab",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
			},
		},
	}
	return armyMap
}
