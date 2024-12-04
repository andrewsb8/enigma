package awgame

type Player struct {
	Id                    int
	Funds                 int
	Units                 []*Unit
	Army_value            int
	Income                int
	Num_income_properties int
	Num_properties        int //including labs and comm towers
}

type Unit struct {
	Id           int
	Type         string
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
	Defense_stars int
	Movement_cost int
	Can_capture   bool
	Capture_value int
	X_location    int
	Y_location    int
	// locations might be redundant. positions will be captured by indices
}

type Map struct {
	Map_width   int
	Map_height  int
	Num_players int
	Has_hq      bool
	Tiles       [][]*Tile
}

type Game struct {
	Awmap          Map
	Num_players    int
	Players        []*Player
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
