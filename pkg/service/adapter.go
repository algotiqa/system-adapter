//=============================================================================
//===
//=== Copyright (C) 2024-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package service

import (
	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/system-adapter/pkg/business"
)

//=============================================================================

func getAdapters(c *auth.Context) {
	list := business.GetAdapters()
	_ = c.ReturnList(list, 0, 1000, len(*list))
}

//=============================================================================

func getAdapter(c *auth.Context) {
	code := c.GetCodeFromUrl()
	a, err := business.GetAdapter(code)

	if err == nil {
		_ = c.ReturnObject(a)
	}

	c.ReturnError(err)
}

//=============================================================================

func getConnectionParams(c *auth.Context) {
	code := c.GetCodeFromUrl()
	config := map[string]any{}
	err := c.BindParamsFromBody(&config)

	if err == nil {
		var params []*adapter.ParamDef
		params, err = business.GetConnectionParams(code, config)
		if err == nil {
			_ = c.ReturnObject(params)
		}
	}

	c.ReturnError(err)
}

//=============================================================================
