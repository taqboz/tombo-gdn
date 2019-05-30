package testdata

import (
	"github.com/taqboz/tombo/cli/config"
)

var Tags, CheckPageParallel, GetLinksParallel, GetLinksTimeSleep = config.ReadJson("testdata/config.json")
