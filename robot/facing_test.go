package robot_test

import (
	"testing"

	"github.com/amir6432/robot/robot"
)

var turnLeftTests = []struct {
	f    robot.Facing
	want robot.Facing
}{
	{robot.East, robot.North},
	{robot.North, robot.West},
	{robot.West, robot.South},
	{robot.South, robot.East},
}

func TestFacing_Left(t *testing.T) {
	for _, tl := range turnLeftTests {
		if got := tl.f.Left(); got != tl.want {
			t.Errorf("f.Left() = %s; want: %s", got, tl.want)
		}
	}
}

var turnRightTests = []struct {
	f    robot.Facing
	want robot.Facing
}{
	{robot.East, robot.South},
	{robot.North, robot.East},
	{robot.West, robot.North},
	{robot.South, robot.West},
}

func TestFacing_Right(t *testing.T) {
	for _, tl := range turnRightTests {
		if got := tl.f.Right(); got != tl.want {
			t.Errorf("f.Right() = %s; want: %s", got, tl.want)
		}
	}
}
