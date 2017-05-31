package robot

// A Facing is one one of the four cardinal directions or points of the compass.
type Facing int

const (
	East Facing = 1 + iota
	North
	West
	South
)

var facings = [...]string{
	"EAST",
	"NORTH",
	"WEST",
	"SOUTH",
}

// String returns the string representation of the Facing
func (f Facing) String() string {
	if East <= f && f <= South {
		return facings[f-1]
	}

	return ""
}

// Left returns the Facing left to the current one (+90 degrees)
func (f Facing) Left() Facing {
	if f < South {
		return f + 1
	}

	return East
}

// Left returns the Facing right to the current one (-90 degrees)
func (f Facing) Right() Facing {
	if East < f {
		return f - 1
	}

	return South
}

// Coefficients returns how the x and y will change when moving in the direction specified by the Facing
func (f Facing) Coefficients() (xco, yco int) {
	switch f {
	case East:
		return +1, 0
	case North:
		return 0, +1
	case West:
		return -1, 0
	case South:
		return 0, -1
	}

	return 0, 0
}
