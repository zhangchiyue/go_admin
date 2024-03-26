package redshift

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (p *Redshift) OnInit() error {
	db, err := gorm.Open(postgres.Open(p.conf.RedShiftConfig.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Logger.LogMode(logger.Error)
	p.DB = db
	return nil
}

func (p *Redshift) Run(closeSig chan bool) {
	<-closeSig
}

func (p *Redshift) stop() error {
	sqlDB, _ := p.DB.DB()
	return sqlDB.Close()
}

func (p *Redshift) OnDestroy() error {
	return p.stop()
}

func (p *Redshift) Name() string {
	return "redshift"
}
