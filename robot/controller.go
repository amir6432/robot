package robot

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Controller parses and dispatches commands to a robot and returns result/error.
// Use NewController to create a controller for a robot.
type Controller interface {
	Command(cmd string) (string, error)
}

// NewController creates and returns a new Controller for the provided robot interface
func NewController(robot Interface) Controller {
	return &controller{robot: robot}
}

type controller struct {
	robot Interface
}

var ErrInvalidCommand = errors.New("invalid command")

var facingStrMap = map[string]Facing{
	"EAST":  East,
	"NORTH": North,
	"WEST":  West,
	"SOUTH": South,
}

// Command parses and dispatches the provided command to the robot
func (c *controller) Command(cmd string) (string, error) {
	cmd = strings.ToUpper(strings.TrimSpace(cmd))

	switch cmd {
	case "LEFT":
		return "", c.robot.Left()
	case "RIGHT":
		return "", c.robot.Right()
	case "MOVE":
		return "", c.robot.Move()
	case "REPORT":
		p, err := c.robot.Report()
		if err != nil {
			return "", err
		}
		return p.String(), nil
	default:
		r, _ := regexp.Compile(`PLACE\s+(\d+)\s*,\s*(\d+)\s*,\s*(EAST|NORTH|WEST|SOUTH)`)
		m := r.FindStringSubmatch(cmd)
		if len(m) == 4 {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])

			return "", c.robot.Place(x, y, facingStrMap[m[3]])
		}
	}

	return "", ErrInvalidCommand
}
