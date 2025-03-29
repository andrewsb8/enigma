# Enigma

This project will be an engine for [Advance Wars By Web](https://awbw.amarriner.com/), akin to Stockfish or equivalent for chess. It is named Enigma after the mission which introduces Sturm in the first Advance Wars game and also because it will utilize machine learning methods to make decisions which means the decision making progress is like a black box and the inner mechanisms are enigmatic.

## Run and Build

From top level directory, run

```$ go run src/enigma.go [options]```

To build, from top level directory, run

```$ go build src/enigma.go```

Which will make a binary in the top level directory which can be run by the command

```$ ./enigma [options]```

## Adding Maps

Training may have to be done on a per-map basis. Map files are stored in ```data/maps/``` and a new directory will be map with the name of the map. You need both the map state, which you can get by downloading a game replay on the relevant map, and the terrain file. The latter you get by choosing ```Export``` on the map info page ([Last Vigil](https://awbw.amarriner.com/prevmaps.php?maps_id=85639)) and copying the numbers ([Terrain IDs](https://awbw.amarriner.com/terrain_map.php)) to a text file in the same location as the map state file.
