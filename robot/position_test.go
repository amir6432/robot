package robot_test

import (
	"testing"

	"github.com/amir6432/robot/robot"
)

func TestPosition_Copy(t *testing.T) {
	p := robot.Position{X: 2, Y: 3, F: robot.East}
	c := p.Copy()

	if p != *c {
		t.Errorf("1Position.Copy() returned a different position %s; want: %s", c, p)
	}

	if &p == c {
		t.Errorf("2Position.Copy() returned the same instance not a copy: %p and %p", c, &p)
	}
}

var positionStringTests = []struct {
	p    robot.Position
	want string
}{
	{robot.Position{1, 3, robot.East}, "1, 3, EAST"},
	{robot.Position{3, 1, robot.West}, "3, 1, WEST"},
	{robot.Position{0, 2, robot.South}, "0, 2, SOUTH"},
	{robot.Position{-1, 1, robot.North}, "-1, 1, NORTH"},
}

func TestPosition_String(t *testing.T) {
	for _, ps := range positionStringTests {
		if got := ps.p.String(); got != ps.want {
			t.Errorf("Postion.String() = %s; want: %s", got, ps.want)
		}
	}
}
