//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package main

import (
	"log/slog"

	"github.com/algotiqa/core/boot"
	"github.com/algotiqa/core/msg"
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/app"
	"github.com/algotiqa/system-adapter/pkg/business"
	"github.com/algotiqa/system-adapter/pkg/process"
	"github.com/algotiqa/system-adapter/pkg/service"
)

//=============================================================================

const component = "system-adapter"

//=============================================================================

func main() {
	cfg := &app.Config{}
	boot.ReadConfig(component, cfg)
	logger := boot.InitLogger(component, &cfg.Application)
	engine := boot.InitEngine(logger, &cfg.Application)
	initClients()
	msg.InitMessaging(&cfg.Messaging)
	service.Init(engine, cfg, logger)
	process.Init(cfg)
	business.Init()
	boot.RunHttpServer(engine, &cfg.Application)
}

//=============================================================================

func initClients() {
	slog.Info("Initializing clients...")
	req.AddDefaultClient("ca.crt", "server.crt", "server.key")
}

//=============================================================================
