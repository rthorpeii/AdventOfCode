// Package day22 has solutions for Day 22 of Advent of Code
// https://adventofcode.com/2020/day/22
package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day22/input.txt"
var testFile string = "day22/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	// fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	// fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// Deck is a playerss
type Deck struct {
	cards      []int
	prevRounds map[string]bool
}

func createDeck(player string) *Deck {
	cards := strings.Split(player, "\n")

	var playerCards []int
	for i := 1; i < len(cards); i++ {
		value, _ := strconv.Atoi(cards[i])
		playerCards = append(playerCards, value)
	}
	prevRounds := make(map[string]bool)
	return &Deck{cards: playerCards, prevRounds: prevRounds}
}

func (deck *Deck) draw() int {
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card
}

func (deck *Deck) saveRound() {
	cards := sliceToString(deck.cards)
	deck.prevRounds[cards] = true
}

func sliceToString(slice []int) string {
	returnStr := ""
	for _, value := range slice {
		returnStr += strconv.Itoa(value) + ","
	}
	return returnStr
}

func (deck Deck) sameRound() bool {
	return deck.prevRounds[sliceToString(deck.cards)]
}

func (deck *Deck) add(newCards []int) {
	deck.cards = append(deck.cards, newCards...)
}

func (deck *Deck) copy(deckLength int) Deck {
	newCards := make([]int, deckLength)
	copy(newCards, deck.cards)
	newPrevRounds := make(map[string]bool)

	return Deck{cards: newCards, prevRounds: newPrevRounds}
}

//Combat is a
type Combat struct {
	p1 Deck
	p2 Deck
}

func (combat *Combat) play() {
	for len(combat.p1.cards) != 0 && len(combat.p2.cards) != 0 {
		p1Card := combat.p1.draw()
		p2Card := combat.p2.draw()

		if p1Card > p2Card {
			combat.p1.add([]int{p1Card, p2Card})
		} else {
			combat.p2.add([]int{p2Card, p1Card})
		}
	}
}

func (combat Combat) sameRound() bool {
	return combat.p1.sameRound() && combat.p2.sameRound()
}

func (combat *Combat) copy(p1Len int, p2Len int) Combat {
	newP1 := combat.p1.copy(p1Len)
	newP2 := combat.p2.copy(p2Len)

	return Combat{newP1, newP2}
}

func (combat *Combat) recursivePlay() int {
	// fmt.Println("\nStarting new Game\n", combat)
	for len(combat.p1.cards) != 0 && len(combat.p2.cards) != 0 {
		// fmt.Println(combat)
		if combat.sameRound() {
			return 1
		}
		combat.p1.saveRound()
		combat.p2.saveRound()

		p1Card := combat.p1.draw()
		p2Card := combat.p2.draw()

		if p1Card <= len(combat.p1.cards) && p2Card <= len(combat.p2.cards) {
			newCombat := combat.copy(p1Card, p2Card)
			winner := newCombat.recursivePlay()
			if winner == 1 {
				combat.p1.add([]int{p1Card, p2Card})
			} else {
				combat.p2.add([]int{p2Card, p1Card})
			}
			continue
		}

		if p1Card > p2Card {
			combat.p1.add([]int{p1Card, p2Card})
		} else {
			combat.p2.add([]int{p2Card, p1Card})
		}
	}

	if len(combat.p1.cards) != 0 {
		return 1
	}
	return 2
}

func (deck Deck) score() int {
	var winningDeck = deck.cards

	score := 0
	for index, card := range winningDeck {
		mult := len(winningDeck) - index
		score += card * mult
	}
	return score
}

func (combat *Combat) score() int {
	var winningDeck []int
	if len(combat.p1.cards) > 0 {
		winningDeck = combat.p1.cards
	} else {
		winningDeck = combat.p2.cards
	}

	score := 0
	for index, card := range winningDeck {
		mult := len(winningDeck) - index
		score += card * mult
	}
	return score
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.ReadInput(file)
	playersInput := strings.Split(rawInput, "\n\n")

	p1Deck := createDeck(playersInput[0])
	p2Deck := createDeck(playersInput[1])

	combat := Combat{*p1Deck, *p2Deck}

	combat.play()
	return combat.score()

}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.ReadInput(file)
	playersInput := strings.Split(rawInput, "\n\n")

	p1Deck := createDeck(playersInput[0])
	p2Deck := createDeck(playersInput[1])

	combat := Combat{*p1Deck, *p2Deck}
	winner := combat.recursivePlay()

	if winner == 1 {
		return combat.p1.score()
	}
	return combat.p2.score()
}
