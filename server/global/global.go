package global

import (
	"context"
	"embed"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	LOG          *log.Logger
	WailsContext *context.Context
	Icon         []byte
	PNGIcon      []byte
	AssetsUI     embed.FS
)
