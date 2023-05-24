package core

import (
	"base-system-backend/global"
	"base-system-backend/initialize"
	"fmt"
)

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	if err := Router.Run(address); err != nil {
		global.LOG.Error(fmt.Errorf("run server failed:  %s", err).Error())
		return
	}
}
