package domain

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// Explorer is our main interface (only implemented by MarsExplorer for now)
type Explorer interface {
	SendInstructions()
}

// MarsBuilder allows us to override / provide a *logrus.logger (should have a particular interface)
type MarsBuilder struct {
	logger *logrus.Logger
}

// Surface is the representation of Mars as a grid
type Surface struct {
	MaxX, MaxY int
}

// Scent is the representation of the trace of a robot which got lost
type Scent struct {
	posX, posY int
	direction  string
}

// MarsExplorer contains all the pieces to execute the instructions to the robots
type MarsExplorer struct {
	Surface *Surface
	Robots  []Robot
	Scents  []Scent
}

// NewMarsBuilder is our MarsBuilder constructor, allows us to get started by injecting a logger
func NewMarsBuilder(logger *logrus.Logger) MarsBuilder {
	return MarsBuilder{logger: logger}
}

// Build is setting up our MarsExplorer
// by receiving a specific array of instructions to setup both our surface and robots
func (mb *MarsBuilder) Build(instructions []string) (*MarsExplorer, error) {
	if len(instructions) == 0 {
		return nil, fmt.Errorf("expected instructions, received %d", len(instructions))
	}

	surface, err := mb.NewSurface(instructions[0])
	if err != nil {
		return nil, fmt.Errorf("failed to build mars surface, got %q", err)
	}

	robots, err := mb.LoadRobotInstructions(instructions[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to load robots instructions, got %q", err)
	}

	return &MarsExplorer{
		Surface: surface,
		Robots:  robots,
	}, nil
}

// SendInstructions is the main entry point to iterate through our robots and execute the given instructions
func (m *MarsExplorer) SendInstructions() {
	for r := range m.Robots {
		if m.isRobotOffBound(m.Robots[r]) {
			// panic / err?
			// continue for now
			continue
		}

		for i := range m.Robots[r].Instructions {
			if m.isThereARobotScent(m.Robots[r], m.Robots[r].Instructions[i]) {
				continue
			}

			// @TODO check for error
			_ = m.Robots[r].Execute(m.Robots[r].Instructions[i])

			if m.isRobotOffBound(m.Robots[r]) {
				m.Robots[r].lost()
				m.leaveScent(m.Robots[r])
				break
			}
		}
	}
}

// NewSurface is the Surface constructor making sure the grid is in order
// The first line of input is the upper-right coordinates of the rectangular world, the lower-left
// coordinates are assumed to be 0, 0.
// The maximum value for any coordinate is 50.
func (mb *MarsBuilder) NewSurface(line string) (*Surface, error) {
	l := strings.Split(line, " ")

	if len(l) != 2 {
		return nil, fmt.Errorf("expected two characters seperated by space, got %d", len(l))
	}

	var maxX int
	maxX, err := strconv.Atoi(l[0])
	if err != nil {
		mb.logger.Errorf(`failed to convert surface X "%s" into integer, got %q`, l[0], err)
		return nil, err
	}
	var maxY int
	maxY, err = strconv.Atoi(l[1])
	if err != nil {
		mb.logger.Errorf(`failed to convert surface Y "%s" into integer, got %q`, l[0], err)
		return nil, err
	}

	if maxX > 50 || maxY > 50 {
		mb.logger.Errorf("maximum value for the grid execeeded 50")
		return nil, fmt.Errorf("maximum value for the grid execeeded 50")
	}

	return &Surface{
		MaxX: maxX,
		MaxY: maxY,
	}, nil
}

// LoadRobotInstructions takes a []string and setup our robots given a specific input
// It consists of a sequence of robot positions and instructions (two lines per
// robot). A position consists of two integers specifying the initial coordinates of the robot and
// an orientation (N, S, E, W), all separated by whitespace on one line. A robot instruction is a
// string of the letters “L”, “R”, and “F” on one line.
func (mb *MarsBuilder) LoadRobotInstructions(lines []string) ([]Robot, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("expected instructions got 0")
	}

	robots := make([]Robot, 0)
	rCount := 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		l := strings.Split(v, " ")
		switch len(l) {
		case 3:
			var posX int
			posX, err := strconv.Atoi(l[0])
			if err != nil {
				mb.logger.Errorf(`failed to convert pos X "%s" into integer, got %q`, l[0], err)
				return nil, err
			}
			var posY int
			posY, err = strconv.Atoi(l[1])
			if err != nil {
				mb.logger.Errorf(`failed to convert pos Y "%s" into integer, got %q`, l[0], err)
				return nil, err
			}
			robots = append(robots, Robot{
				PosX:      posX,
				PosY:      posY,
				Direction: l[2], // @TODO validate direction or maybe on execution
			})
			continue
		case 1:
			// @TODO validate instructions  or maybe on execution
			robots[rCount].Instructions = strings.SplitAfter(v, "")
			rCount++
			continue
		default:
			continue
		}
	}

	return robots, nil
}

// isRobotOffBound asserts a robot is still on the planet
func (m *MarsExplorer) isRobotOffBound(r Robot) bool {
	if r.PosY > m.Surface.MaxY {
		return true
	}

	if r.PosX > m.Surface.MaxX {
		return true
	}

	return false
}

// isThereARobotScent verify if there isn't a robot's scent left for that grid position
func (m *MarsExplorer) isThereARobotScent(r Robot, c string) bool {
	if len(m.Scents) == 0 {
		return false
	}

	for _, s := range m.Scents {
		// sounds like there might a better way to do this
		if s.posY == r.PosY && s.posX == r.PosX && s.direction == r.Direction && c == CommandForward {
			return true
		}
	}

	return false
}

// leaveScent create a new scent when a robot got lost
func (m *MarsExplorer) leaveScent(r Robot) {
	m.Scents = append(m.Scents, Scent{
		posX:      r.PosX,
		posY:      r.PosY,
		direction: r.Direction,
	})
}
