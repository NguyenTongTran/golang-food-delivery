package appctx

import "gorm.io/gorm"

type ctx struct {
	mainDB *gorm.DB
}

func New(mainDB *gorm.DB) *ctx {
	return &ctx{mainDB}
}

func (c ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}
