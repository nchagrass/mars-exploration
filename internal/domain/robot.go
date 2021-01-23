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

// Robot representation of a robot
type Robot struct {
	PosX         int
	PosY         int
	Direction    string
	Instructions []string
	Lost         bool
}

// ControlledRobot available commands to execute on a Robot
type ControlledRobot interface {
	Execute(command string) error
	forward() error
	turnRight() error
	turnLeft() error
	isLost() bool
}

// Execute will execute the corresponding command on a robot if the instruction is recognised
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

// turnRight moves the robot in a different direction in a clockwise manner
// if the robot has an unrecognised direction it will return an error
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

// turnLeft moves the robot in a different direction in a counterclockwise manner
// if the robot has an unrecognised direction it will return an error
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

// forward moves the robot on the grid via increasing/decreasing it's Y or X position
// if the robot has an unrecognised direction it will return an error
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

// backward reverse the robot to its previous position on the grid
// if the robot has an unrecognised direction it will return an error
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

// lost marks a robot as lost (used when gone out of the grid)
func (r *Robot) lost() {
	r.Lost = true
	// panic / err up
	_ = r.backward()
}

// isLost returns the status of a Robot
func (r *Robot) isLost() bool {
	return r.Lost
}

// ToString returns a pre-defined output as a string for a report of the robot status
func (r *Robot) ToString() string {
	if r.isLost() {
		return fmt.Sprintf("%d %d %s %s", r.PosX, r.PosY, r.Direction, "LOST")
	}

	return fmt.Sprintf("%d %d %s", r.PosX, r.PosY, r.Direction)
}
