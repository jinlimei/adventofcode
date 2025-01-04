package day06

type direction uint8

const (
	facingUnknown direction = iota
	facingNorth
	facingEast
	facingSouth
	facingWest
)

func (gd direction) String() string {
	switch gd {
	case facingNorth:
		return "^"
	case facingEast:
		return ">"
	case facingSouth:
		return "v"
	case facingWest:
		return "<"
	default:
		panic("invalid direction")
	}
}

func (gd direction) turnRight() direction {
	switch gd {
	case facingNorth:
		return facingEast
	case facingEast:
		return facingSouth
	case facingSouth:
		return facingWest
	case facingWest:
		return facingNorth
	default:
		panic("invalid direction")
	}
}
