package bootstrap

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

type Robot struct {
	PosX         int
	PosY		 int
	Direction    string
	Instructions *[]string
}

func (m *MarsExplorer) ExecuteInstructions() {
	// @TODO
}

// Bootstrap initialise the project
func New() (Explorer, error) {
	// load configuration
	// logger / output

	// parse given file path

	// fall back on a default input path

	// load mars grid

	// load robots instructions

	// exec

	return nil, nil
}
