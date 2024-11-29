package awio

import (
	"log"
	"os"
	"strings"
)

func GetMapState(map_file string) string {
	/*
	 Reads map file from AWBW and returns the final line
	 of data as a string to be parsed.
	*/
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
