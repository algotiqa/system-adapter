//=============================================================================
//===
//=== Copyright (C) 2024-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package business

import (
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/system-adapter/pkg/adapter/local"
	"github.com/algotiqa/system-adapter/pkg/adapter/tradestation"
)

//=============================================================================

var adapters map[string]adapter.Adapter
var infos []*adapter.Info

//=============================================================================
//===
//=== Init
//===
//=============================================================================

func init() {
	adapters = map[string]adapter.Adapter{}

	register(local.NewAdapter())
	register(tradestation.NewAdapter())
	//	register(interactive .NewAdapter())
}

//=============================================================================

func register(a adapter.Adapter) {
	info := a.GetInfo()
	adapters[info.Code] = a
	infos = append(infos, info)
}

//=============================================================================
//===
//=== Public methods
//===
//=============================================================================

func GetAdapters() *[]*adapter.Info {
	return &infos
}

//=============================================================================

func GetAdapter(code string) (*adapter.Info, error) {
	a, ok := adapters[code]
	if !ok {
		return nil, req.NewNotFoundError(code)
	}

	return a.GetInfo(), nil
}

//=============================================================================

func GetConnectionParams(code string, configParams map[string]any) ([]*adapter.ParamDef, error) {
	a, ok := adapters[code]
	if !ok {
		return nil, req.NewNotFoundError(code)
	}

	return a.GetConnectParams(configParams), nil
}

//=============================================================================
