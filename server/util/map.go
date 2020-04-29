package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethanmil/bolo/guide"
)

// BuildMapFromFile -
func BuildMapFromFile() *guide.WorldMap {
	wm := &guide.WorldMap{}
	file, err := os.Open("assets/test_map.txt")
	if err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}

	wm.SizeW = int32(len(lines[0])/2 + 1)
	wm.SizeH = int32(len(lines))

	wm.Tiles = make([]string, wm.SizeH*wm.SizeW)
	for y := 0; y < int(wm.SizeH); y++ {
		for x, tileType := range strings.Split(lines[y], ",") {
			seq := x + (y * int(wm.SizeH))
			wm.Tiles[seq] = tileType
		}
	}

	return wm
}
