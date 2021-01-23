package bootstrap

import (
	"github.com/sirupsen/logrus"
	"marsrobot/internal/domain"
)

// Bootstrap initialise the project
func New(path string) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	setup, err := NewFileInstructions(path)
	if err != nil {
		logger.Fatalf(`unable to read instructions from path "%s" - got %q`, path, err)
	}

	// load mars grid / robots
	builder := domain.NewMarsBuilder(logger)
	me, err := builder.Build(setup)
	if err != nil {
		logrus.Fatalf("failed to prepare the exploration, %q", err)
	}

	me.SendInstructions()

	reporter := domain.Reporter{Explorer: me}

	reporter.Print()
}
