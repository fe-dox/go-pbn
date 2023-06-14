package pbn

import (
	"strconv"
	"strings"
)

type BoardSet struct {
	EventName string
	Generator string
	Boards    []Board
}

type Hand map[Suit][]CardValue

func (h *Hand) String() string {
	var b strings.Builder
	var i int
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for _, card := range (*h)[suit] {
			b.WriteString(card.String())
		}
		i++
		if i != 4 {
			b.WriteString(".")
		}
	}
	return b.String()
}

func NewHand() Hand {
	return map[Suit][]CardValue{
		Spades:   make([]CardValue, 0, 5),
		Hearts:   make([]CardValue, 0, 5),
		Diamonds: make([]CardValue, 0, 5),
		Clubs:    make([]CardValue, 0, 5),
	}
}

type Board struct {
	Number       int
	Dealer       Direction
	Vulnerable   Vulnerability
	EventName    string
	Generator    string
	Hands        map[Direction]Hand
	Ability      Ability
	OptimumScore struct {
		Direction Direction
		Score     int
	}
	MinimaxScore Contract
}

func (m *Contract) String() string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(m.Level))
	sb.WriteString(m.Suit.ShortString())

	if m.Doubled {
		sb.WriteString("D")
	}

	sb.WriteString(m.Direction.String())
	sb.WriteString(strconv.Itoa(m.Score))

	return sb.String()
}

func NewBoard() *Board {
	return &Board{
		Number:     0,
		Dealer:     0,
		Vulnerable: 0,
		Hands: map[Direction]Hand{
			North: NewHand(),
			East:  NewHand(),
			South: NewHand(),
			West:  NewHand(),
		},
	}
}

type Contract struct {
	Level     int
	Suit      Suit
	Doubled   bool
	Redoubled bool
	Direction Direction
	Score     int
}

type Ability map[Direction]map[Suit]int

func (a *Ability) String() string {
	var sb strings.Builder
	var i int

	for _, direction := range []Direction{North, East, South, West} {
		sb.WriteString(direction.String())
		sb.WriteString(":")
		results := (*a)[direction]
		for _, suit := range []Suit{NoTrump, Spades, Hearts, Diamonds, Clubs} {
			result := results[suit]
			if result < 10 {
				sb.WriteString(strconv.Itoa(result))
			} else {
				switch result {
				case 10:
					sb.WriteString("A")
				case 11:
					sb.WriteString("B")
				case 12:
					sb.WriteString("C")
				case 13:
					sb.WriteString("D")
				}
			}
		}

		i++
		if i != 4 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

func NewOptimumResultTable() *Ability {
	return &Ability{
		North: map[Suit]int{
			NoTrump:  0,
			Spades:   0,
			Hearts:   0,
			Diamonds: 0,
			Clubs:    0,
		},
		East: map[Suit]int{
			NoTrump:  0,
			Spades:   0,
			Hearts:   0,
			Diamonds: 0,
			Clubs:    0,
		},
		South: map[Suit]int{
			NoTrump:  0,
			Spades:   0,
			Hearts:   0,
			Diamonds: 0,
			Clubs:    0,
		},
		West: map[Suit]int{
			NoTrump:  0,
			Spades:   0,
			Hearts:   0,
			Diamonds: 0,
			Clubs:    0,
		},
	}
}
