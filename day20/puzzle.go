// Package day20 has solutions for Day 20 of Advent of Code
// https://adventofcode.com/2020/day/20
package day20

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day20/input.txt"
var testFile string = "day20/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// Tile represents a tile of the image
type Tile struct {
	id    int
	data  []string
	edges []string
}

func (tile Tile) leftEdge() string {
	left := ""
	for _, line := range tile.data {
		left += string(line[0])
	}
	return left
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
func (tile Tile) bottomEdge() string {
	return tile.data[len(tile.data)-1]
}

// Rotates a tile 90 degrees clockwise
func (tile *Tile) rotate() {
	newData := make([]string, len(tile.data))
	for x := 0; x < len(tile.data[0]); x++ {
		newLine := ""
		for y := len(tile.data) - 1; y >= 0; y-- {
			newLine += string(tile.data[y][x])
		}
		newData[x] = newLine
	}
	tile.data = newData
}

func newTile(rawTile string) Tile {
	split := strings.Fields(rawTile)
	id, _ := strconv.Atoi(split[1][:len(split[1])-1])
	return Tile{
		id:   id,
		data: split[2:]}
}

func (tile *Tile) print() {
	fmt.Println("ID: ", tile.id)
	for _, line := range tile.data {
		fmt.Println(line)
	}
}

func getStartingTiles(file string) map[int]*Tile {
	rawInput := strings.Split(input.ReadInput(file), "\n\n")

	tiles := make(map[int]*Tile, len(rawInput))
	for _, rawTile := range rawInput {
		tile := newTile(rawTile)
		tiles[tile.id] = &tile
	}

	return tiles
}

func generateEdgeMap(tiles *map[int]*Tile) map[string]map[int]bool {
	edgeMap := make(map[string]map[int]bool)

	for id, tile := range *tiles {
		edges := tile.getEdges()
		for _, edge := range edges {
			_, exists := edgeMap[edge]
			if !exists {
				edgeMap[edge] = make(map[int]bool)
			}
			edgeMap[edge][id] = true
		}
	}

	return edgeMap
}

func (tile *Tile) getEdges() []string {
	if len(tile.edges) != 0 {
		return tile.edges
	}
	edges := tile.getEdgesHelper()
	flippedTile := tile.getFlipped()
	edges = append(edges, (flippedTile.getEdgesHelper())...)

	tile.edges = edges
	return tile.edges
}

func (tile *Tile) getEdgesHelper() []string {
	var edges []string
	right := ""
	left := ""
	for _, line := range tile.data {
		right += string(line[len(line)-1])
		left = string(line[0]) + left
	}
	bottom := Reverse(tile.data[len(tile.data)-1])
	return append(edges, tile.data[0], right, left, bottom)
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (tile *Tile) flipTile() {
	newData := make([]string, len(tile.data))

	for index, line := range tile.data {
		newIndex := len(tile.data) - index - 1
		newData[newIndex] = line
	}
	tile.data = newData
}

func (tile Tile) getFlipped() Tile {
	newData := make([]string, len(tile.data))

	for index, line := range tile.data {
		newIndex := len(tile.data) - index - 1
		newData[newIndex] = line
	}

	return Tile{
		id:    tile.id,
		data:  newData,
		edges: tile.edges,
	}
}

func findCorners(borderIDs *map[int]bool, allTiles *map[int]*Tile, edgeMap *map[string]map[int]bool) []Tile {
	var borderTiles []Tile
	for id := range *borderIDs {
		borderCount := 0
		tile := (*allTiles)[id]
		for _, edge := range tile.edges {
			if len((*edgeMap)[edge]) == 1 {
				borderCount++
			}
		}
		if borderCount > 2 {
			borderTiles = append(borderTiles, *tile)
		}
	}
	return borderTiles
}

// PartOne finds
func PartOne(file string) int {
	tiles := getStartingTiles(file)
	edgeMap := generateEdgeMap(&tiles)

	borderIDs := make(map[int]bool)
	for _, values := range edgeMap {
		for id := range values {
			if len(values) == 1 {
				borderIDs[id] = true
			}
		}
	}

	corners := findCorners(&borderIDs, &tiles, &edgeMap)
	fmt.Println(len(corners), corners)

	value := 1
	for _, tile := range corners {
		value *= tile.id
	}
	return value
}

// PartTwo finds
func PartTwo(file string) int {
	tiles := getStartingTiles(file)
	edgeMap := generateEdgeMap(&tiles)

	borderIDs := make(map[int]bool)
	for _, values := range edgeMap {
		for id := range values {
			if len(values) == 1 {
				borderIDs[id] = true
			}
		}
	}

	// corners := findCorners(&borderIDs, &tiles, &edgeMap)

	puzzle := Puzzle{
		tiles:     tiles,
		edgeMap:   edgeMap,
		borderIDs: borderIDs,
	}
	corners := findCorners(&borderIDs, &tiles, &edgeMap)

	puzzle.solve(corners[0])
	puzzle.printAssembled()

	image := combineTiles(puzzle.assembled)
	fmt.Println("IMAGE: ")
	for _, line := range image {
		fmt.Println(line)
	}

	imageTile := Tile{data: image, id: -1}

	findMonsters(&imageTile)
	imageTile.print()

	count := 0
	for _, line := range imageTile.data {
		for _, char := range line {
			if string(char) == "#" {
				count++
			}
		}
	}

	return count
}

func findMonsters(image *Tile) {
	for i := 0; i < 3; i++ {
		matched := matchMonsters(&image.data)
		if matched {
			return
		}
		image.rotate()
	}
	matched := matchMonsters(&image.data)
	if !matched {
		image.flipTile()
		findMonsters(image)
	}
}

func matchMonsters(image *[]string) bool {
	sawMonster := false
	for y := 0; y < len(*image)-len(seaMonster); y++ {
	IteratePixel:
		for x := 0; x < len((*image)[y])-len(seaMonster[0]); x++ {
			for dy := 0; dy < len(seaMonster); dy++ {
				for dx := 0; dx < len(seaMonster[0]); dx++ {
					imageX := x + dx
					imageY := y + dy

					if seaMonster[dy][dx] == '#' {
						if string((*image)[imageY][imageX]) != "#" {
							continue IteratePixel
						}
					}
				}
			}
			sawMonster = true
			for dy := 0; dy < len(seaMonster); dy++ {
				for dx := 0; dx < len(seaMonster[0]); dx++ {
					imageX := x + dx
					imageY := y + dy

					if seaMonster[dy][dx] == '#' {
						(*image)[imageY] = (*image)[imageY][:imageX] + "O" + (*image)[imageY][imageX+1:]
					}
				}
			}
		}
	}
	return sawMonster
}

var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   "}

func combineTiles(tiles [][]Tile) []string {
	var image []string

	var strippedTiles [][][]string

	for _, line := range tiles {
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
			image = append(image, imageLine)
		}
	}
	return image
}

func (tile Tile) stripBorder() []string {
	lines := make([]string, len(tile.data)-2)

	for y := 1; y < len(tile.data)-1; y++ {
		line := ""
		for x := 1; x < len(tile.data[y])-1; x++ {
			line += string(tile.data[y][x])
		}
		lines[y-1] = line
	}
	return lines
}

func (puzzle Puzzle) printAssembled() {
	tileLength := len(puzzle.assembled[0][0].data)
	for _, line := range puzzle.assembled {
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

// Puzzle represents the entire puzzle
type Puzzle struct {
	tiles     map[int]*Tile
	edgeMap   map[string]map[int]bool
	borderIDs map[int]bool
	assembled [][]Tile
}

func (puzzle *Puzzle) solve(corner Tile) {
	puzzle.orientCorner(&corner)
	borderLength := int(math.Sqrt(float64(len(puzzle.tiles))))
	assembled := make([][]Tile, borderLength)
	for y := 0; y < borderLength; y++ {
		row := make([]Tile, borderLength)
		for x := 0; x < borderLength; x++ {
			if x == 0 { // We're along the left edge
				if y == 0 { // We're in the top left corner
					row[0] = corner
					continue
				}

				// We need to match on the row above us
				prevTile := assembled[y-1][x]
				sideToMatch := prevTile.bottomEdge()
				nextID := puzzle.findMatchingID(prevTile.id, sideToMatch)

				nextTile := puzzle.tiles[nextID]
				nextTile.orientTop(sideToMatch)
				row[x] = *nextTile
				continue
			}
			// Match on the tile to our left
			prevTile := row[x-1]
			sideToMatch := prevTile.rightEdge()
			nextID := puzzle.findMatchingID(prevTile.id, sideToMatch)
			nextTile := puzzle.tiles[nextID]

			nextTile.orientLeft(sideToMatch)
			row[x] = *nextTile
		}
		assembled[y] = row
	}
	puzzle.assembled = assembled
}

func (puzzle Puzzle) findMatchingID(tileID int, edge string) int {
	for id := range puzzle.edgeMap[edge] {
		if id != tileID {
			return id
		}
	}
	log.Fatal("Couldn't find matching tile")
	return -1
}

func (puzzle Puzzle) orientCorner(corner *Tile) {
	for i := 0; i < 3; i++ {
		left := corner.leftEdge()
		top := corner.topEdge()
		if (len(puzzle.edgeMap[left]) == 1) && (len(puzzle.edgeMap[top]) == 1) {
			return
		}
		corner.rotate()
	}
}

func (tile *Tile) orientLeft(rightEdge string) {
	for i := 0; i < 3; i++ {
		left := tile.leftEdge()
		if left == rightEdge {
			return
		}
		tile.rotate()
	}
	left := tile.leftEdge()
	if left != rightEdge { // If it didn't align flip it and try again
		tile.flipTile()
		tile.orientLeft(rightEdge)
	}
}

func (tile *Tile) orientTop(bottomEdge string) {
	for i := 0; i < 3; i++ {
		top := tile.topEdge()
		if top == bottomEdge {
			return
		}
		tile.rotate()
	}
	top := tile.topEdge()
	if top != bottomEdge { // If it didn't align flip it and try again
		tile.flipTile()
		tile.orientTop(bottomEdge)
	}
}
