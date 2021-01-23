package domain

import "fmt"

const (
	CommandRight   = "R"
	CommandLeft    = "L"
	CommandForward = "F"

	DirectionNorth = "N"
	DirectionEast  = "E"
	DirectionSouth = "S"
	DirectionWest  = "W"
)

type Robot struct {
	PosX         int
	PosY         int
	Direction    string
	Instructions []string
	Lost         bool
}

type ControlledRobot interface {
	Execute(command string) error
	forward() error
	turnRight() error
	turnLeft() error
	isLost() bool
}

func (r *Robot) Execute(c string) error {
	switch c {
	case CommandRight:
		return r.turnRight()
	case CommandLeft:
		return r.turnLeft()
	case CommandForward:
		return r.forward()
	default:
		return fmt.Errorf("unsupported command: %s", c)
	}
}

func (r *Robot) turnRight() error {
	// @TODO maybe come back and play with a list of directions
	switch r.Direction {
	case DirectionEast:
		r.Direction = DirectionSouth
		return nil
	case DirectionSouth:
		r.Direction = DirectionWest
		return nil
	case DirectionWest:
		r.Direction = DirectionNorth
		return nil
	case DirectionNorth:
		r.Direction = DirectionEast
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}

func (r *Robot) turnLeft() error {
	switch r.Direction {
	case DirectionEast:
		r.Direction = DirectionNorth
		return nil
	case DirectionNorth:
		r.Direction = DirectionWest
		return nil
	case DirectionWest:
		r.Direction = DirectionSouth
		return nil
	case DirectionSouth:
		r.Direction = DirectionEast
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}

func (r *Robot) forward() error {
	switch r.Direction {
	case DirectionEast:
		r.PosX = r.PosX + 1
		return nil
	case DirectionWest:
		r.PosX = r.PosX - 1
		return nil
	case DirectionNorth:
		r.PosY = r.PosY + 1
		return nil
	case DirectionSouth:
		r.PosY = r.PosY - 1
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}

func (r *Robot) backward() error {
	switch r.Direction {
	case DirectionEast:
		r.PosX = r.PosX - 1
		return nil
	case DirectionWest:
		r.PosX = r.PosX + 1
		return nil
	case DirectionNorth:
		r.PosY = r.PosY - 1
		return nil
	case DirectionSouth:
		r.PosY = r.PosY + 1
		return nil
	default:
		return fmt.Errorf("unsupported Robot direction %s", r.Direction)
	}
}

func (r *Robot) lost() {
	r.Lost = true
	// panic / err up
	_ = r.backward()
}

func (r *Robot) isLost() bool {
	return r.Lost
}

func (r *Robot) ToString() string {
	if r.isLost() {
		return fmt.Sprintf("%d %d %s %s", r.PosX, r.PosY, r.Direction, "LOST")
	}

	return fmt.Sprintf("%d %d %s", r.PosX, r.PosY, r.Direction)
}
