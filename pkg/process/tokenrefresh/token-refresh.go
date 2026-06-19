//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package tokenrefresh

import (
	"log/slog"
	"time"

	"github.com/algotiqa/core/msg"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/system-adapter/pkg/app"
	"github.com/algotiqa/system-adapter/pkg/business"
)

//=============================================================================

func InitRefresh(cfg *app.Config) *time.Ticker {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		run()

		for range ticker.C {
			run()
		}
	}()

	return ticker
}

//=============================================================================

func run() {
	list := business.GetConnectionsToRefresh()

	for _, ctx := range list {
		err := ctx.RefreshToken()
		if err != nil {
			slog.Error("TokenRefresher: Cannot refresh token. Disconnecting", "username", ctx.Username, "connection", ctx.ConnectionCode, "error", err.Error())
			err = sendConnectionChangeMessage(ctx)
			if err != nil {
				slog.Error("TokenRefresher:  Could not publish the disconnection message (!)", "username", ctx.Username, "connection", ctx.ConnectionCode, "error", err.Error())
			}
		} else {
			slog.Info("TokenRefresher: Refreshed token complete", "username", ctx.Username, "connection", ctx.ConnectionCode)
		}
	}
}

//=============================================================================

func sendConnectionChangeMessage(ctx *adapter.ConnectionContext) error {
	ccm := business.ConnectionChangeSystemMessage{
		Username:       ctx.Username,
		ConnectionCode: ctx.ConnectionCode,
		SystemCode:     ctx.GetAdapterInfo().Code,
		Status:         ctx.GetStatus(),
	}

	return msg.SendMessage(msg.ExSystem, msg.SourceConnection, msg.TypeChange, &ccm, nil)
}

//=============================================================================
