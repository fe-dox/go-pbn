package pbn

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

func DirectionFromRune(r rune) Direction {
	switch r {
	case 'N', 'n':
		return North
	case 'E', 'e':
		return East
	case 'S', 's':
		return South
	case 'W', 'w':
		return West
	default:
		return North
	}
}

func SideFromDirection(direction Direction) string {
	switch direction {
	case North:
		return "NS"
	case South:
		return "NS"
	case West:
		return "EW"
	case East:
		return "EW"
	default:
		return "NS"
	}
}

func DirectionFromString(s string) Direction {
	return DirectionFromRune([]rune(s)[0])
}

//goland:noinspection GoUnusedExportedFunction
func DealerFromBoardNumber(n int) Direction {
	switch n % 4 {
	case 1:
		return North
	case 2:
		return East
	case 3:
		return South
	case 0:
		return West
	default:
		return North
	}
}

type Vulnerability int

const (
	None Vulnerability = iota
	NorthSouth
	EastWest
	Both
)

func (v Vulnerability) String() string {
	return [...]string{"None", "NS", "EW", "All"}[v]
}

func VulnerabilityFromString(s string) Vulnerability {
	switch s {
	case "None", "Love", "none", "love":
		return None
	case "NS", "ns":
		return NorthSouth
	case "EW", "ew":
		return EastWest
	case "Both", "All", "all", "both":
		return Both
	default:
		return None
	}
}

//goland:noinspection GoUnusedExportedFunction
func VulnerabilityFromBoardNumber(n int) Vulnerability {
	switch n % 4 {
	case 1:
		return None
	case 2:
		return NorthSouth
	case 3:
		return EastWest
	case 0:
		return Both
	default:
		return None
	}
}

type Suit int

const (
	NoTrump Suit = iota
	Spades
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	return [...]string{"NT", "S", "H", "D", "C"}[s]
}

func (s Suit) ShortString() string {
	return [...]string{"N", "S", "H", "D", "C"}[s]
}

func SuitFromSting(s string) Suit {
	switch s {
	case "NT", "nt", "N", "n":
		return NoTrump
	case "S", "s":
		return Spades
	case "H", "h":
		return Hearts
	case "D", "d":
		return Diamonds
	case "C", "c":
		return Clubs
	default:
		return NoTrump
	}
}

type CardValue int

//goland:noinspection GoUnusedConst
const (
	UndefinedCard CardValue = 0
	A             CardValue = 1
	T             CardValue = 10
	J             CardValue = 11
	Q             CardValue = 12
	K             CardValue = 13
)

func (c CardValue) String() string {
	return [...]string{"U", "A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}[c]
}

func CardValueFromRune(r rune) CardValue {
	switch r {
	case 'A':
		return A
	case 'T':
		return T
	case 'J':
		return J
	case 'Q':
		return Q
	case 'K':
		return K
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return CardValue(r - '0')
	default:
		return CardValue(0)
	}
}

//goland:noinspection GoUnusedExportedFunction
func CardValueFromString(s string) CardValue {
	return CardValueFromRune([]rune(s)[0])
}
