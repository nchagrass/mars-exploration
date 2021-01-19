package bootstrap

import (
	"github.com/sirupsen/logrus"
	"marsrobot/internal/domain"
)

// Bootstrap initialise the project
func New(path string) (domain.Explorer, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	_, err := NewFileInstructions(path)
	if err != nil {
		logger.Errorf(`unable to read instructions from path "%s" - got %q`, path, err)
		return nil, err
	}

	// load mars grid

	// load robots instructions

	// exec

	return nil, nil
}
