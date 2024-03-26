package redshift

import (
	"adx-admin/pkg/configer"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

type Redshift struct {
	DB   *gorm.DB
	conf *configer.Config
}

func NewDatabase(conf *configer.Config) *Redshift {
	var instance *Redshift
	once.Do(func() {
		instance = &Redshift{
			DB:   nil,
			conf: conf,
		}
	})
	return instance
}
