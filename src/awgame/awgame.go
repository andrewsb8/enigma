package awgame

type Player struct {
	Id                    int
	Country_id            int
	CO_id                 int
	Funds                 int
	Units                 []*Unit
	Building_IDs          []int //ids of captured buildings
	Army_value            int
	Income                int
	Num_income_properties int
	Num_properties        int //including labs and comm towers
}

type Unit struct {
	// Type_id is id in state/action file
	Unit_id  int
	Type     string
	Movement int // total movement
	// index of movment cost array to search
	// when computing cost of movement paths
	Movement_type int
	Vision        int
	X_position    int
	Y_position    int
	Value         int
	Hit_points    int
	Ammo          int
	Fuel          int
	Fuel_per_turn int
	Can_capture   bool
	Is_capturing  bool
}

type Tile struct {
	Terrain_id    int
	Terrain_type  string
	Defense_stars int
	// Unit order of below array will follow the wiki:
	// [ foot boot treads tires sea lander air pipe ]
	// if unit cannot move on terrain, gets value of 100
	Movement_cost_clear [8]int
	Movement_cost_rain  [8]int
	Movement_cost_snow  [8]int
	Can_capture         bool
	Capture_points      int
	// id from state file which will help keep track of
	// which player has which building
	Building_id int
}

type Map struct {
	Map_width   int
	Map_height  int
	Num_players int
	Has_hq      bool
	// indices of Tiles are the coordinates of the map
	Tiles [][]*Tile
}

type Game struct {
	Awmap          Map
	Num_players    int
	Players        []*Player
	Country_ids    []*int //list of country ids for easy searching
	Day            int
	Fog            bool
	Starting_funds int
	Weather        string
}

func ParseGameInfo(map_state string, terrain []string) *Game {
	game := Game{}
	ParseTerrain(terrain, &game)
	ParseMapState(map_state, &game)
	return &game
}
