package main

import (
  "os"
  "fmt"
  "flag"

  p "enigma/src/play"
)

func usage() {
  fmt.Fprintf(os.Stderr, "usage: enigma [options] [values]\n")
  flag.PrintDefaults()
  os.Exit(0)
}

var (
  play = flag.Bool("p", false, "Flag to tell enigma to play a game.")
  input_map = flag.String("im", "", "Input file with map information")
)

func main() {
  flag.Usage = usage
  flag.Parse()

  if *play {
    p.Play()
  } else {
    usage()
  }
}
