package awio

import (
	"log"
	"os"
	"strings"
)

/*
Reads map file from AWBW and returns the final line
of data as a string to be parsed.
*/
func GetMapState(map_file string) string {
	data, err := os.ReadFile(map_file)
	if err != nil {
		log.Fatal(err)
	} else {
		lines := strings.Split(string(data), "\n")
		//-2 because these files have an empty line at the end
		return lines[len(lines)-2]
	}
	return "" //would like to remove this
}

func GetTerrain(terrain_file string) []string {
	data, err := os.ReadFile(terrain_file)
	if err != nil {
		log.Fatal(err)
	} else {
		return strings.Split(string(data), "\n")
	}
	return []string{} //would like to remove this
}
