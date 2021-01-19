package domain

type Explorer interface {
	ExecuteInstructions()
}

type Surface struct {
	MaxX, MaxY int
}

type MarsExplorer struct {
	Surface *Surface
	Robots []Robot
}

func NewMarsExplorer(instructions []string) *MarsExplorer {
	// load surface

	// load robot

	return &MarsExplorer{}
}

func (m *MarsExplorer) ExecuteInstructions() {
	// @TODO
}