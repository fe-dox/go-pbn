package pbn

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParsePBN(s io.Reader) *BoardSet {
	result := BoardSet{}
	scanner := bufio.NewScanner(s)
	currentBoard := NewBoard()
	expectOptimumResultTable := false
	var currentOptimumResultTable *Ability
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "%") || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "{") {
			continue
		}
		if len(line) == 0 && currentBoard != nil {
			if expectOptimumResultTable {
				expectOptimumResultTable = false
				if currentOptimumResultTable != nil {
					currentBoard.Ability = *currentOptimumResultTable
					currentOptimumResultTable = nil
				}
			}
			result.Boards = append(result.Boards, *currentBoard)
			currentBoard = NewBoard()
			continue
		}
		if strings.HasPrefix(line, "[") {
			var tag string
			if expectOptimumResultTable {
				expectOptimumResultTable = false
				if currentOptimumResultTable != nil {
					currentBoard.Ability = *currentOptimumResultTable
					currentOptimumResultTable = nil
				}
			}
			_, err := fmt.Sscanf(line, "[%s ", &tag)
			if err != nil {
				fmt.Println(err)
			}
			switch tag {
			case "Event":
				eventName := strings.TrimSuffix(strings.TrimPrefix(line, "[Event \""), "\"]")
				result.EventName = eventName
				currentBoard.EventName = eventName
			case "Generator":
				generator := strings.TrimSuffix(strings.TrimPrefix(line, "[Generator \""), "\"]")
				result.Generator = generator
				currentBoard.Generator = generator
			case "Board":
				boardNumber, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(line, "[Board \""), "\"]"))
				if err != nil {
					fmt.Println(err)
				}
				currentBoard.Number = boardNumber
			case "Dealer":
				dealer := strings.TrimSuffix(strings.TrimPrefix(line, "[Dealer \""), "\"]")
				switch dealer {
				case "N", "n":
					currentBoard.Dealer = North
				case "E", "e":
					currentBoard.Dealer = East
				case "S", "s":
					currentBoard.Dealer = South
				case "W", "w":
					currentBoard.Dealer = West
				}
			case "Vulnerable":
				rawVulnerable := strings.TrimSuffix(strings.TrimPrefix(line, "[Vulnerable \""), "\"]")
				currentBoard.Vulnerable = VulnerabilityFromString(rawVulnerable)
			case "Deal":
				rawDeal := strings.TrimSuffix(strings.TrimPrefix(line, "[Deal \""), "\"]")
				currentDirection := North
				currentHand := NewHand()
				currentColor := Spades
				for _, char := range rawDeal {
					switch char {
					case ' ':
						currentBoard.Hands[currentDirection] = currentHand
						currentDirection++
						currentHand = Hand{}
						currentColor = Spades
					case '.':
						currentColor++
					case '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A':
						currentHand[currentColor] = append(currentHand[currentColor], CardValueFromRune(char))
					}
				}
				currentBoard.Hands[currentDirection] = currentHand
			case "Ability":
				rawAbility := strings.TrimSuffix(strings.TrimPrefix(line, "[Ability \""), "\"]")
				currentDirection := North
				currentSuit := NoTrump
				ability := map[Direction]map[Suit]int{
					North: {},
					East:  {},
					South: {},
					West:  {},
				}
				for _, r := range rawAbility {
					switch r {
					case 'N', 'E', 'S', 'W':
						currentDirection = DirectionFromRune(r)
						currentSuit = NoTrump
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						ability[currentDirection][currentSuit], _ = strconv.Atoi(string(r))
						currentSuit++
					case 'A', 'B', 'C', 'D':
						v := 0
						switch r {
						case 'A':
							v = 10
						case 'B':
							v = 11
						case 'C':
							v = 12
						case 'D':
							v = 13
						}
						ability[currentDirection][currentSuit] = v
						currentSuit++
					}
				}
				currentBoard.Ability = ability
			case "OptimumResultTable":
				expectOptimumResultTable = true
				currentOptimumResultTable = NewOptimumResultTable()
			case "OptimumScore":
				rawOptimumScore := strings.Split(strings.TrimSuffix(strings.TrimPrefix(line, "[OptimumScore \""), "\"]"), " ")
				if len(rawOptimumScore) != 2 {
					continue
				}
				currentBoard.OptimumScore.Direction = DirectionFromString(rawOptimumScore[0])
				currentBoard.OptimumScore.Score, err = strconv.Atoi(rawOptimumScore[1])
				if err != nil {
					fmt.Println(err)
				}
			case "Minimax":
				rawMinimax := strings.Split(strings.TrimPrefix(strings.TrimSuffix(line, "\"]"), "[Minimax \""), "")
				currentBoard.MinimaxScore.Level, err = strconv.Atoi(rawMinimax[0])
				currentBoard.MinimaxScore.Suit = SuitFromSting(rawMinimax[1])
				directionIndex := 2
				if rawMinimax[2] == "D" || rawMinimax[2] == "d" {
					currentBoard.MinimaxScore.Doubled = true
					directionIndex = 3
				}
				currentBoard.MinimaxScore.Direction = DirectionFromString(rawMinimax[directionIndex])
				currentBoard.MinimaxScore.Score, _ = strconv.Atoi(strings.Join(rawMinimax[directionIndex+1:], ""))

			default:
				fmt.Println("Unsupported tag: ", tag)
			}
		} else if expectOptimumResultTable {
			rawOptimumRecord := strings.Split(line, " ")
			if len(rawOptimumRecord) != 3 {
				fmt.Println(rawOptimumRecord)
			}
			(*currentOptimumResultTable)[DirectionFromString(rawOptimumRecord[0])][SuitFromSting(rawOptimumRecord[1])], _ = strconv.Atoi(rawOptimumRecord[2])
		}
	}

	return &result
}
