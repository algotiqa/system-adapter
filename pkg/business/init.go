//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package business

import (
	"log/slog"
	"os"

	"github.com/algotiqa/core/msg"
)

//=============================================================================

func Init() {
	sendSystemRestartMessage()
}

//=============================================================================
//===
//=== Private methods
//===
//=============================================================================

func sendSystemRestartMessage() {
	err := msg.SendMessage(msg.ExSystem, msg.SourceSystem, msg.TypeRestart, nil, nil)

	if err != nil {
		slog.Error("sendSystemRestartMessage: Could not publish the restart message", "error", err.Error())
		os.Exit(1)
	}
}

//=============================================================================
