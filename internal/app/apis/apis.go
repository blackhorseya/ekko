package apis

import "github.com/google/wire"

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	HealthSet,
)
