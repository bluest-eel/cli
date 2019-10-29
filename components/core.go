package components

import (
	"github.com/bluest-eel/cli/components/config"
	"github.com/geomyidia/reverb"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// Base component collection
type Base struct {
	Config *config.Config
	Logger *logger.Logger
}

// TestBase component that keeps stdout clean
type TestBase struct {
	Config *config.Config
	suite.Suite
}

// Default component collection
type Default struct {
	Base
	GRPCD *reverb.Reverb
}

// CLI component collection
type CLI struct {
	Base
	GRPCD *reverb.Reverb
}

// TestGRPC ...
type TestGRPC struct {
	Config *config.Config
	GRPCD  *reverb.Reverb
}

// Add more components here that have more or less than what's done above. This
// is useful for testing or runnning in different binaries/executables, etc.
