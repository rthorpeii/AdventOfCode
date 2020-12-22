package day20

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/helper"
	"github.com/rthorpeii/AdventOfCode2020/input"
)

// Image represents the entire Image
type Image struct {
	tiles      map[int]*Tile
	edgeMap    map[string]map[int]bool
	layout     [][]Tile
	finalImage []string
}

func newImage(file string) *Image {
	rawTiles := strings.Split(input.ReadInput(file), "\n\n")
	tiles := make(map[int]*Tile, len(rawTiles))
	for _, rawTile := range rawTiles {
		tile := newTile(rawTile)
		tiles[tile.id] = &tile
	}

	edgeMap := make(map[string]map[int]bool)
	for id, tile := range tiles {
		edges := tile.getEdges()
		for _, edge := range edges {
			_, exists := edgeMap[edge]
			if !exists {
				edgeMap[edge] = make(map[int]bool)
			}
			edgeMap[edge][id] = true
		}
	}

	return &Image{
		tiles:   tiles,
		edgeMap: edgeMap,
	}
}

func (image Image) printAssembled() {
	tileLength := len(image.layout[0][0].data)
	for _, line := range image.layout {
		lineToPrint := ""
		for i := 0; i < tileLength; i++ {
			for _, tile := range line {
				lineToPrint += tile.data[i] + " "
			}
			lineToPrint += "\n"
		}
		fmt.Println(lineToPrint)
	}
}

/*
Methods for finding the correct layout of the tiles and assembling the final image
*/
func (image *Image) assemble() {
	borderLength := int(math.Sqrt(float64(len(image.tiles))))
	image.layout = make([][]Tile, borderLength)
	for y := 0; y < borderLength; y++ {
		image.layout[y] = make([]Tile, borderLength)
		for x := 0; x < borderLength; x++ {
			// Set the top left corner
			if x == 0 && y == 0 {
				corner := image.findCorners()[0]
				image.orientCorner(&corner)
				image.layout[y][x] = corner
				continue
			}
			image.layout[y][x] = image.getNextTile(x, y)
		}
	}
}

func (image *Image) findCorners() []Tile {

	borderIDs := make(map[int]bool)
	for _, values := range image.edgeMap {
		for id := range values {
			if len(values) == 1 {
				borderIDs[id] = true
			}
		}
	}

	var borderTiles []Tile
	for id := range borderIDs {
		borderCount := 0
		tile := image.tiles[id]
		for _, edge := range tile.edges {
			if len(image.edgeMap[edge]) == 1 {
				borderCount++
			}
		}
		if borderCount > 2 {
			borderTiles = append(borderTiles, *tile)
		}
	}
	return borderTiles
}

func (image *Image) getNextTile(x, y int) Tile {
	var prevTile Tile
	var edgeToMatch string
	var edgeFunc func(Tile) string
	if x == 0 {
		// If we're along the left edge, match on the tile above us
		prevTile = image.layout[y-1][x]
		edgeToMatch = prevTile.bottomEdge()
		edgeFunc = topEdge
	} else {
		// Match on the tile to our left
		prevTile = image.layout[y][x-1]
		edgeToMatch = prevTile.rightEdge()
		edgeFunc = leftEdge
	}

	nextID := image.findMatchingID(prevTile.id, edgeToMatch)
	nextTile := image.tiles[nextID]
	nextTile.orientEdge(edgeToMatch, edgeFunc)
	return *nextTile
}

func (image Image) findMatchingID(tileID int, edge string) int {
	for id := range image.edgeMap[edge] {
		if id != tileID {
			return id
		}
	}
	log.Fatal("Couldn't find matching tile")
	return -1
}

// rotates a corner piece until it can be placed in the top left of the image
func (image Image) orientCorner(corner *Tile) {
	for i := 0; i < 3; i++ {
		left, top := corner.leftEdge(), corner.topEdge()
		if (len(image.edgeMap[left]) == 1) && (len(image.edgeMap[top]) == 1) {
			return
		}
		corner.rotate()
	}
}

func (image *Image) createFinalImage() {
	var strippedTiles [][][]string

	for _, line := range image.layout {
		var strippedLine [][]string
		for _, tile := range line {
			strippedLine = append(strippedLine, tile.stripBorder())
		}
		strippedTiles = append(strippedTiles, strippedLine)
	}

	for _, line := range strippedTiles {
		for i := 0; i < len(line[0]); i++ {
			imageLine := ""
			for _, tile := range line {
				imageLine += tile[i]
			}
			image.finalImage = append(image.finalImage, imageLine)
		}
	}
}

/*
Methods for finding monsters within the image
*/
func (image *Image) findMonsters(monster *[]string) {
	for i := 0; i < 3; i++ {
		matched := image.markMonsters(monster)
		if matched {
			return
		}
		image.finalImage = helper.RotateSlice(&image.finalImage)
	}
	matched := image.markMonsters(monster)
	if !matched {
		image.finalImage = helper.FlipSlice(&image.finalImage)
		image.findMonsters(monster)
	}
}

func (image *Image) markMonsters(seaMonster *[]string) bool {
	sawMonster := false
	for y := 0; y < len(image.finalImage)-len(*seaMonster); y++ {
		for x := 0; x < len(image.finalImage[y])-len((*seaMonster)[0]); x++ {
			if !image.matchMonster(x, y, seaMonster) {
				continue
			}
			sawMonster = true
			image.markMonster(x, y, seaMonster)
		}
	}
	return sawMonster
}

func (image *Image) matchMonster(x, y int, seaMonster *[]string) bool {
	for dy := 0; dy < len(*seaMonster); dy++ {
		for dx := 0; dx < len((*seaMonster)[0]); dx++ {
			if (*seaMonster)[dy][dx] == '#' && (*image).finalImage[y+dy][x+dx] != '#' {
				return false
			}
		}
	}
	return true
}

func (image *Image) markMonster(x, y int, seaMonster *[]string) {
	for dy := 0; dy < len(*seaMonster); dy++ {
		for dx := 0; dx < len((*seaMonster)[0]); dx++ {
			imageX := x + dx
			imageY := y + dy

			if (*seaMonster)[dy][dx] == '#' {
				image.finalImage[imageY] = image.finalImage[imageY][:imageX] + "O" + image.finalImage[imageY][imageX+1:]
			}
		}
	}
}
