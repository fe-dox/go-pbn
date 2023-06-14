package pbn

import (
	"fmt"
	"io"
	"log"
)

func (bs *BoardSet) Serialize(w io.Writer, abilityAsTable bool) error {
	useBoardSetMetadata := bs.EventName != "" || bs.Generator != ""
	for i := 0; i < len(bs.Boards); i++ {
		if useBoardSetMetadata {
			bs.Boards[i].EventName = bs.EventName
			bs.Boards[i].Generator = bs.Generator
		}
		err := bs.Boards[i].Serialize(w, abilityAsTable)
		if err != nil {
			log.Printf("Error occured during parsing board number %d (%d): %s ", i, bs.Boards[i].Number, err.Error())
			return err
		}
	}
	return nil
}

func (b *Board) Serialize(w io.Writer, abilityAsTable bool) error {
	var err error
	if b.Generator != "" {
		err = WriteTag("Generator", b.Generator, w)
		if err != nil {
			return err
		}
	}
	if b.EventName != "" {
		err = WriteTag("Event", b.EventName, w)
	}
	if b.Number != 0 {
		err = WriteTag("Board", fmt.Sprintf("%d", b.Number), w)
		if err != nil {
			return err
		}
	}
	err = WriteTag("Dealer", b.Dealer.String(), w)
	if err != nil {
		return err
	}
	err = WriteTag("Vulnerable", b.Vulnerable.String(), w)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, "[Deal \"N:")
	if err != nil {
		return err
	}
	var j int
	for _, direction := range []Direction{North, East, South, West} {
		hand := b.Hands[direction]
		_, err = io.WriteString(w, hand.String())
		j++
		if j != 4 {
			_, err = io.WriteString(w, " ")
		}
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, "\"]\n")
	if err != nil {
		return err
	}
	if b.Ability != nil {
		if abilityAsTable {
			err := WriteOptimumResultTable(w, b.Ability)
			if err != nil {
				return err
			}
		}
		err = WriteTag("Ability", b.Ability.String(), w)
		if err != nil {
			return err
		}
	}

	err = WriteTag("OptimumScore", fmt.Sprintf("%s %d", SideFromDirection(b.OptimumScore.Direction), b.OptimumScore.Score), w)
	if err != nil {
		return err
	}
	if b.MinimaxScore.Level != 0 {
		err = WriteTag("Minimax", b.MinimaxScore.String(), w)
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, "\n")
	if err != nil {
		return err
	}
	return nil

}

func WriteTag(tagName string, value string, w io.Writer) error {
	_, err := io.WriteString(w, fmt.Sprintf("[%s \"%s\"]\n", tagName, value))
	if err != nil {
		return err
	}
	return nil
}

func WriteOptimumResultTable(w io.Writer, ability Ability) error {
	err := WriteTag("OptimumResultTable", "Declarer;Denomination\\2R;Result\\2R", w)
	if err != nil {
		return err
	}
	for _, direction := range []Direction{North, East, South, West} {
		for _, suit := range []Suit{Clubs, Diamonds, Hearts, Spades, NoTrump} {
			_, err = io.WriteString(w, fmt.Sprintf("%s %s %d\n", direction.String(), suit.String(), ability[direction][suit]))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
