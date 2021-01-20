package domain

import "fmt"

type Robot struct {
	PosX         int
	PosY         int
	Direction    string
	Instructions []string
}

type ControlledRobot interface {
	Execute(command string) error
}

func (r *Robot) Execute(c string) error {
	switch c {
	case "R": // turn right
		return r.turnRight()

	case "L": // turn left
		return r.turnLeft()
	case "F":
		return nil
	default:
		return fmt.Errorf("unsupported command: %s", c)
	}
}

func (r *Robot) turnRight() error {
	// @TODO maybe come back and play with a list of directions
	switch r.Direction {
	case "E":
		r.Direction = "S"
		return nil
	case "S":
		r.Direction = "W"
		return nil
	case "W":
		r.Direction = "N"
		return nil
	case "N":
		r.Direction = "E"
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}

func (r *Robot) turnLeft() error {
	switch r.Direction {
	case "E":
		r.Direction = "N"
		return nil
	case "N":
		r.Direction = "W"
		return nil
	case "W":
		r.Direction = "S"
		return nil
	case "S":
		r.Direction = "E"
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}
