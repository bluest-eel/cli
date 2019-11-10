package components

import (
	"github.com/bluest-eel/cli/components/config"
	"github.com/bluest-eel/common/components"
	// "google.golang.org/grpc"
)

// Base component collection
type Base struct {
	Config *config.Config
	components.BaseLogger
}

// BaseApp ...
type BaseApp struct {
	components.BaseApp
}

// BaseCLI ...
type BaseCLI struct {
	components.BaseCLI
}

// Default component collection
type Default struct {
	Base
	components.BaseGRPC
}

// TestBase component that keeps stdout clean
type TestBase struct {
	Config *config.Config
	components.TestBase
}

// TestGRPC ...
type TestGRPC struct {
	Config *config.Config
	components.TestBaseGRPC
}
