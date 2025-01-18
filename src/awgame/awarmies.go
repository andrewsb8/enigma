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
				Can_capture:         true,
				Capture_points:      20,
			},
			54: {
				Terrain_type:        "base",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			55: {
				Terrain_type:        "airport",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			56: {
				Terrain_type:        "port",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			57: {
				Terrain_type:        "hq",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			136: {
				Terrain_type:        "comm tower",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			148: {
				Terrain_type:        "lab",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
		},
		20: { //purple lightning
			172: {
				Terrain_type:        "city",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			171: {
				Terrain_type:        "base",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			170: {
				Terrain_type:        "airport",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			176: {
				Terrain_type:        "port",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			174: {
				Terrain_type:        "hq",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			173: {
				Terrain_type:        "comm tower",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
			175: {
				Terrain_type:        "lab",
				Defense_stars:       1,
				Movement_cost_clear: [8]int{1, 1, 1, 2, 100, 100, 1, 100},
				Can_capture:         true,
				Capture_points:      20,
			},
		},
	}
	return armyMap
}
