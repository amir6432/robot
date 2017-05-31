package robot

import "errors"

// Interface is an interface for a robot
type Interface interface {

	// Place puts the robot at the provided position and facing
	// An error is returned if the place is out of table's bounds
	Place(x int, y int, f Facing) error

	// Place puts the robot at the provided position and facing
	// An error is returned if the place is out of table's bounds
	PlacePosition(p Position) error

	// Left turns robot's direction +90 degrees
	// An error is returned if the robot is not placed on a table yet.
	Left() error

	// Left turns robot's direction -90 degrees
	// An error is returned if the robot is not placed on a table yet.
	Right() error

	// Move moves the robot one step in its current facing
	// An error is returned if the robot is not placed on a table yet
	// or if the move will put the robot out of table's bounds.
	Move() error

	// Report returns the current position of the robot.
	// An error is returned if the robot is not placed on a table yet.
	Report() (*Position, error)
}

var ErrOutOfBounds = errors.New("robot cannot go out of bounds")
var ErrNotPlaced = errors.New("robot is not placed on the table")

// New creates a new robot Interface connected to the provided table
func New(table Table) Interface {
	return &robot{table: table}
}

type robot struct {
	table    Table
	position *Position
}

func (r *robot) Place(x int, y int, f Facing) error {
	if !r.table.InBounds(x, y) {
		return ErrOutOfBounds
	}

	if r.position == nil {
		r.position = &Position{}
	}

	r.position.X, r.position.Y, r.position.F = x, y, f

	return nil
}

func (r *robot) PlacePosition(p Position) error {
	return r.Place(p.X, p.Y, p.F)
}

func (r *robot) Left() error {
	if r.position == nil {
		return ErrNotPlaced
	}

	r.position.F = r.position.F.Left()
	return nil
}

func (r *robot) Right() error {
	if r.position == nil {
		return ErrNotPlaced
	}

	r.position.F = r.position.F.Right()
	return nil
}

func (r *robot) Move() error {
	if r.position == nil {
		return ErrNotPlaced
	}

	cx, cy := r.position.F.Coefficients()
	newX, newY := r.position.X+cx, r.position.Y+cy

	if !r.table.InBounds(newX, newY) {
		return ErrOutOfBounds
	}

	r.position.X, r.position.Y = newX, newY
	return nil
}

func (r *robot) Report() (*Position, error) {
	if r.position == nil {
		return nil, ErrNotPlaced
	}

	return r.position.Copy(), nil
}
