package main

import (
	"flag"
	"fmt"
	"os"

	p "enigma/src/play"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: enigma [options] [values]\n")
	flag.PrintDefaults()
	os.Exit(0)
}

var (
	play         = flag.Bool("p", false, "Flag to tell enigma to play a game.")
	map_file     = flag.String("mf", "", "Input file with map state information")
	action_file  = flag.String("af", "", "Input file with action information")
	terrain_file = flag.String("tf", "", "Input file with grid of terrain ids")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *play {
		p.Play(*map_file, *terrain_file)
	} else {
		usage()
	}
}
