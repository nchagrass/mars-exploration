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

	// load robots

	return &MarsExplorer{
		Surface: surface,
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
		mb.logger.Errorf(`failed to convert surface X "%s" into integer(s), got %q`, l[0], err)
		return nil, err
	}
	var maxY int
	maxY, err = strconv.Atoi(l[1])
	if err != nil {
		mb.logger.Errorf(`failed to convert surface Y "%s" into integer(s), got %q`, l[0], err)
		return nil, err
	}

	return &Surface{
		MaxX: maxX,
		MaxY: maxY,
	}, nil

}
