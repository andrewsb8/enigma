package main

import (
  "os"
  "fmt"
  "flag"
)

func usage() {
  fmt.Fprintf(os.Stderr, "usage: enigma [options] [values]\n")
  flag.PrintDefaults()
  os.Exit(0)
}

var (
  play = flag.Bool("p", false, "Flag to tell enigma to play a game.")
)

func main() {
  flag.Usage = usage
  flag.Parse()

  if *play {
    fmt.Printf("cool\n")
  } else {
    usage()
  }
}
