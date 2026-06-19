//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package business

import (
	"sync"

	"github.com/algotiqa/system-adapter/pkg/adapter"
)

//=============================================================================

type ConnectionSpec struct {
	SystemCode    string         `json:"systemCode"     binding:"required"`
	ConfigValues  map[string]any `json:"configValues"   binding:"required"`
	ConnectValues map[string]any `json:"connectValues"  binding:"required"`
}

//=============================================================================

type ConnectionInfo struct {
	Username       string `json:"username"`
	ConnectionCode string `json:"connectionCode"`
	SystemCode     string `json:"systemCode"`
	SystemName     string `json:"systemName"`
	Status         string `json:"status"`
}

//=============================================================================

type UserConnections struct {
	sync.RWMutex
	username string
	contexts map[string]*adapter.ConnectionContext
}

//-----------------------------------------------------------------------------

func NewUserConnections() *UserConnections {
	uc := &UserConnections{}
	uc.contexts = make(map[string]*adapter.ConnectionContext)

	return uc
}

//=============================================================================

type ConnectionChangeSystemMessage struct {
	Username       string                `json:"username"`
	ConnectionCode string                `json:"connectionCode"`
	SystemCode     string                `json:"systemCode"`
	Status         adapter.ContextStatus `json:"status"`
}

//=============================================================================

type TestAdapterRequest struct {
	Service string `json:"service"`
	Query   string `json:"query"`
}

//=============================================================================
