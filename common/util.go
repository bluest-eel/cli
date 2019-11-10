package common

import (
	"runtime"

	"github.com/bluest-eel/common/util"
)

// CallerPaths see the util.CallerPaths function for details
func CallerPaths() *util.CallerTree {
	return util.CallerPaths(runtime.Caller(0))
}
