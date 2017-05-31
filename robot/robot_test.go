package robot_test

import (
	"testing"

	"github.com/amir6432/robot/robot"
)

type mockTable struct {
	inBounds bool
}

func (t *mockTable) InBounds(x, y int) bool {
	return t.inBounds
}

func TestRobot_Left_shouldFailIfNotPlaced(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	if err := r.Left(); err != robot.ErrNotPlaced {
		t.Error("Robot.Left() should fail when robot is not on the table")
	}
}

func TestRobot_Right_shouldFailIfNotPlaced(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	if err := r.Right(); err != robot.ErrNotPlaced {
		t.Error("Robot.Right() should fail when robot is not on the table")
	}
}

func TestRobot_Move_shouldFailIfNotPlaced(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	if err := r.Move(); err != robot.ErrNotPlaced {
		t.Error("Robot.Move() should fail when robot is not on the table")
	}
}

func TestRobot_Report_shouldFailIfNotPlaced(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	if _, err := r.Report(); err != robot.ErrNotPlaced {
		t.Error("Robot.Report() should fail when robot is not on the table")
	}
}

func TestRobot_Place_shouldFailIfOutOfBounds(t *testing.T) {
	mt := &mockTable{inBounds: false}
	r := robot.New(mt)

	r.Place(1, 1, robot.East)
	if err := r.Place(1, 1, robot.East); err != robot.ErrOutOfBounds {
		t.Error("Robot.Place() should fail if out of bounds")
	}
}

func TestRobot_Move_shouldFailIfOutOfBounds(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	r.Place(1, 1, robot.East)

	mt.inBounds = false
	if err := r.Move(); err != robot.ErrOutOfBounds {
		t.Error("Robot.Move() should fail if out of bounds")
	}
}

func TestRobot_Left(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	r.Place(1, 1, robot.East)
	r.Left()

	p, _ := r.Report()
	if p.F != robot.North {
		t.Errorf("Robot.Left() failed to change facing. got = %s; want: %s", p.F, robot.North)
	}
}

func TestRobot_Right(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	r.Place(1, 1, robot.East)
	r.Right()

	p, _ := r.Report()
	if p.F != robot.South {
		t.Errorf("Robot.Right() failed to change facing. got = %s; want: %s", p.F, robot.South)
	}
}

func TestRobot_Report(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	p := robot.Position{X: 1, Y: 1, F: robot.East}
	r.Place(p.X, p.Y, p.F)

	pr, _ := r.Report()
	if p != *pr {
		t.Errorf("Robot.Report() = %s; want: %s", pr, p)
	}
}

func TestRobot_Report_shouldNotBeAbleToChangePosition(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	p := robot.Position{X: 1, Y: 1, F: robot.East}
	r.PlacePosition(p)

	pr, _ := r.Report()
	pr.X, pr.Y, pr.F = 3, 4, robot.South

	pr2, _ := r.Report()
	if *pr2 != p {
		t.Errorf("Robot.Report() = %s; want: %s", pr2, p)
	}
}

var moveTests = []struct {
	start robot.Position
	want  robot.Position
}{
	{robot.Position{1, 1, robot.East}, robot.Position{2, 1, robot.East}},
	{robot.Position{0, 0, robot.East}, robot.Position{1, 0, robot.East}},
	{robot.Position{3, 5, robot.North}, robot.Position{3, 6, robot.North}},
	{robot.Position{2, 4, robot.West}, robot.Position{1, 4, robot.West}},
	{robot.Position{2, 4, robot.South}, robot.Position{2, 3, robot.South}},
}

func TestRobot_Move(t *testing.T) {
	mt := &mockTable{inBounds: true}
	r := robot.New(mt)

	for _, mt := range moveTests {
		r.PlacePosition(mt.start)
		r.Move()
		pr, _ := r.Report()
		if *pr != mt.want {
			t.Errorf("Robot.Move() from %s resulted in %s; want: %s", mt.start, pr, mt.want)
		}
	}
}
