package database

import (
	"adx-admin/pkg/database/rds"
	"adx-admin/pkg/database/redshift"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(redshift.NewDatabase, rds.NewRedisDB)
