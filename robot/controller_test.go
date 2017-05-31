package robot_test

import (
	"fmt"
	"testing"

	"github.com/amir6432/robot/robot"
)

type mockRobot struct {
	commands string
}

func (r *mockRobot) Place(x int, y int, f robot.Facing) error {
	r.commands += fmt.Sprintf("Place%d,%d,%s", x, y, f)

	return nil
}

func (r *mockRobot) PlacePosition(p robot.Position) error {
	return r.Place(p.X, p.Y, p.F)
}

func (r *mockRobot) Left() error {
	r.commands += "Left"

	return nil
}

func (r *mockRobot) Right() error {
	r.commands += "Right"

	return nil
}

func (r *mockRobot) Move() error {
	r.commands += "Move"

	return nil
}

func (r *mockRobot) Report() (*robot.Position, error) {
	r.commands += "Report"

	return &robot.Position{X: 1, Y: 1, F: robot.North}, nil
}

var commandTests = []struct {
	commands []string
	want     string
	errors   int
	output   string
}{
	{
		[]string{"LEFT", "MOVE", "REPORT"},
		"LeftMoveReport",
		0,
		"1, 1, NORTH",
	},
	{
		[]string{"PLACE 2, 3, NORTH", "RIGHT"},
		"Place2,3,NORTHRight",
		0,
		"",
	},
	{
		[]string{"left", "move ", " rePort"}, //accept lower case and white spaces
		"LeftMoveReport",
		0,
		"1, 1, NORTH",
	},
	{
		[]string{"LEFT", "IGNORE", "aaaaa", "REPORT"}, //skip invalid commands
		"LeftReport",
		2,
		"1, 1, NORTH",
	},
	{
		[]string{"PLACE 2, 3", "RIGHT"}, //skip incomplete PLACE command
		"Right",
		1,
		"",
	},
}

func TestController_Command(t *testing.T) {
	for _, ct := range commandTests {
		r := &mockRobot{}
		c := robot.NewController(r)

		errCount := 0
		output := ""
		for _, cmd := range ct.commands {
			o, err := c.Command(cmd)
			output += o
			if err != nil {
				errCount++
			}
		}
		if ct.want != r.commands {
			t.Errorf("c.Command() dispatched %s; want: %s", r.commands, ct.want)
		}
		if ct.errors != errCount {
			t.Errorf("c.Command() returned %d errors; want: %d", errCount, ct.errors)
		}
		if ct.output != output {
			t.Errorf("c.Command() produced %s; want: %s", output, ct.output)
		}
	}
}
