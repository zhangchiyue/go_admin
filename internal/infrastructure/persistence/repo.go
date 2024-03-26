package persistence

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRShiftReportRepo,
)
