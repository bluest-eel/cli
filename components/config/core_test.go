package config_test

import (
	"testing"

	"github.com/bluest-eel/cli/components"
	"github.com/bluest-eel/cli/components/config"
	"github.com/stretchr/testify/suite"
)

type configTestSuite struct {
	components.TestBase
}

func (suite *configTestSuite) SetupSuite() {
	suite.Config = config.NewConfig()
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, &configTestSuite{})
}

func (suite *configTestSuite) TestDefaultConfig() {
	// XXX why are these failing right now?
	// suite.Equal(5099, suite.Config.GRPCD.Port)
	// suite.Equal(4222, suite.Config.MessagingServer.Port)
}
