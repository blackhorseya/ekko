package config

import "github.com/google/wire"

// ProviderSet is a configs provider set
var ProviderSet = wire.NewSet(NewConfig)
