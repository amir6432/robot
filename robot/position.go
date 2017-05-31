package robot

import (
	"fmt"
)

type Position struct {
	X int
	Y int
	F Facing
}

func (p Position) String() string {
	return fmt.Sprintf("%d, %d, %s", p.X, p.Y, p.F)
}

func (p *Position) Copy() *Position {
	clone := *p
	return &clone
}
