//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package service

import (
	"log/slog"

	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/core/auth/roles"
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/app"
	"github.com/gin-gonic/gin"
)

//=============================================================================

func Init(router *gin.Engine, cfg *app.Config, logger *slog.Logger) {

	ctrl := auth.NewOidcController(cfg.Authentication.Authority, req.GetDefaultClient(), logger, cfg)

	//--- Adapter services

	router.GET ("/api/system/v1/adapters",                  ctrl.Secure(getAdapters, roles.Admin_User))
	router.GET ("/api/system/v1/adapters/:code",            ctrl.Secure(getAdapter, roles.Admin_User))
	router.POST("/api/system/v1/adapters/:code/connection", ctrl.Secure(getConnectionParams, roles.Admin_User))

	//--- Connection services

	router.GET   ("/api/system/v1/connections",                                ctrl.Secure(getConnections, roles.Admin_User))
	router.PUT   ("/api/system/v1/connections/:code",                          ctrl.Secure(connect,        roles.Admin_User))
	router.DELETE("/api/system/v1/connections/:code",                          ctrl.Secure(disconnect,     roles.Admin_User))
	router.GET   ("/api/system/v1/connections/:code/roots",                    ctrl.Secure(getRootSymbols, roles.Admin_User))
	router.GET   ("/api/system/v1/connections/:code/roots/:root",              ctrl.Secure(getRootSymbol,  roles.Admin_User))
	router.GET   ("/api/system/v1/connections/:code/roots/:root/instruments",  ctrl.Secure(getInstruments, roles.Admin_User_Service))
	router.GET   ("/api/system/v1/connections/:code/instruments/:symbol/bars", ctrl.Secure(getPriceBars,   roles.Admin_User_Service))
	router.GET   ("/api/system/v1/connections/:code/accounts",                 ctrl.Secure(getAccounts,    roles.Admin_User_Service))
	router.GET   ("/api/system/v1/connections/:code/orders",                   ctrl.Secure(getOrders,      roles.Admin_User_Service))
	router.GET   ("/api/system/v1/connections/:code/positions",                ctrl.Secure(getPositions,   roles.Admin_User_Service))
	router.POST  ("/api/system/v1/connections/:code/test",                     ctrl.Secure(testAdapter,    roles.Admin_User))

	router.GET("/api/system/v1/connections/:code/login", webLogin)
	router.Use(proxyLoginRequests)
}

//=============================================================================
