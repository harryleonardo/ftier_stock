package context

import (
	SharedConfig "github.com/ftier-stock/shared/config"
	"github.com/labstack/echo"
)

// ApplicationContext ...
type ApplicationContext struct {
	echo.Context
	Config SharedConfig.ImmutableConfig
}
