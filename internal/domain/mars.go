package domain

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Explorer interface {
	ExecuteInstructions()
}

type MarsBuilder struct {
	logger *logrus.Logger
}

type Surface struct {
	MaxX, MaxY int
}

type MarsExplorer struct {
	Surface *Surface
	Robots  []Robot
}

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

func (m *MarsExplorer) ExecuteInstructions() {
	// @TODO
}

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

	return &Surface{
		MaxX: maxX,
		MaxY: maxY,
	}, nil
}

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
				Direction: l[2], // @TODO validate direction
			})
			continue
		case 1:
			// @TODO validate instructions
			robots[rCount].Instructions = strings.SplitAfter(v, "")
			rCount++
			continue
		default:
			continue
		}
	}

	return robots, nil
}
