package awio

import (
	"fmt"
	"os"
	"strings"
)

func ReadMap(map_file string) {
	data, err := os.ReadFile(map_file)
	if err != nil {
		fmt.Printf("Error reading map file: %s\n", err)
	} else {
		lines := strings.Split(string(data), "\n")
		//fmt.Println(lines[0])
		split_by_curly := strings.Split(lines[1], "}")
		fmt.Println(split_by_curly[1])
	}
}
