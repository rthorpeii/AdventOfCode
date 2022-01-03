package day20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/helper"
)

// Tile represents a small tile that is part of a larger image
type Tile struct {
	id    int
	data  []string
	edges []string
}

// newTile creates a new tile from a string representation of a tile
func newTile(rawTile string) Tile {
	split := strings.Fields(rawTile)
	id, _ := strconv.Atoi(split[1][:len(split[1])-1])
	return Tile{
		id:   id,
		data: split[2:]}
}

func (tile Tile) print() {
	fmt.Println("ID: ", tile.id)
	for _, line := range tile.data {
		fmt.Println(line)
	}
}

// Strips the border text of a tile
func (tile Tile) stripBorder() []string {
	lines := make([]string, len(tile.data)-2)
	for y := 1; y < len(tile.data)-1; y++ {
		lines[y-1] = tile.data[y][1 : len(tile.data[y])-1]
	}
	return lines
}

/*
Helper methods related to getting the edges of a tile
*/
func (tile Tile) leftEdge() string {
	left := ""
	for _, line := range tile.data {
		left += string(line[0])
	}
	return left
}
func leftEdge(tile Tile) string {
	return tile.leftEdge()
}

func (tile Tile) rightEdge() string {
	right := ""
	for _, line := range tile.data {
		right += string(line[len(line)-1])
	}
	return right
}

func (tile Tile) topEdge() string {
	return tile.data[0]
}
func topEdge(tile Tile) string {
	return tile.topEdge()
}

func (tile Tile) bottomEdge() string {
	return tile.data[len(tile.data)-1]
}

func (tile *Tile) getEdges() []string {
	if len(tile.edges) != 0 {
		return tile.edges
	}
	edges := tile.getEdgesHelper()
	tile.flipTile()
	edges = append(edges, (tile.getEdgesHelper())...)

	tile.edges = edges
	return tile.edges
}
func (tile *Tile) getEdgesHelper() []string {
	var edges []string
	right, left := "", ""
	for _, line := range tile.data {
		right += string(line[len(line)-1])
		left = string(line[0]) + left
	}
	bottom := helper.Reverse(tile.data[len(tile.data)-1])
	return append(edges, tile.data[0], right, left, bottom)
}

/*
Helper methods related to orienting tiles
*/

// Rotates a tile 90 degrees clockwise
func (tile *Tile) rotate() {
	tile.data = helper.RotateSlice(&tile.data)
}

// flipTile flips a tile along the horizontal axis
func (tile *Tile) flipTile() {
	tile.data = helper.FlipSlice(&tile.data)
}

// orientEdge manipulates the facing of a tile until the edge returned by edgeFunc matches
// the specified edge
func (tile *Tile) orientEdge(edgeToMatch string, edgeFunc func(tile Tile) string) {
	targetEdge := edgeFunc(*tile)
	// Rotate the tile until the specified edge matches the target
	for rotations := 0; rotations < 3; rotations++ {
		if targetEdge == edgeToMatch {
			return
		}
		tile.rotate()
		targetEdge = edgeFunc(*tile)
	}
	// If it didn't align flip it and try again
	if targetEdge != edgeToMatch {
		tile.flipTile()
		tile.orientEdge(edgeToMatch, edgeFunc)
	}
}
