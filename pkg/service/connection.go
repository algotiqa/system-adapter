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
	"net/http"
	"net/url"

	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/system-adapter/pkg/business"
	"github.com/algotiqa/types"
	"github.com/gin-gonic/gin"
)

//=============================================================================

func getConnections(c *auth.Context) {
	var filter map[string]any
	offset, limit, err := c.GetPagingParams()

	if err == nil {
		list := business.GetConnections(c, filter, offset, limit)
		_ = c.ReturnList(list, 0, 1000, len(*list))
	}

	c.ReturnError(err)
}

//=============================================================================

func connect(c *auth.Context) {
	code := c.GetCodeFromUrl()
	spec := business.ConnectionSpec{}
	err := c.BindParamsFromBody(&spec)

	if err == nil {
		var res *adapter.ConnectionResult
		res, err = business.Connect(c, code, &spec)
		if err == nil {
			_ = c.ReturnObject(res)
			return
		}
	}

	c.ReturnError(err)
}

//=============================================================================

func disconnect(c *auth.Context) {
	code := c.GetCodeFromUrl()
	err := business.Disconnect(c, code)

	if err == nil {
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func webLogin(c *gin.Context) {
	code := c.Param("code")

	ctx := business.GetConnectionContextByInstanceCode(code)
	if ctx == nil {
		req.ReturnError(c, req.NewBadRequestError("Connection context not found : "+code))
		return
	}

	authUrl := ctx.GetAdapterAuthUrl()
	target, err := url.Parse(authUrl)
	if err != nil {
		req.ReturnError(c, req.NewBadRequestError("Bad authentication url : "+authUrl))
		return
	}

	c.SetCookie(InstanceCode, code, 0, "", "", true, true)

	location := target.Path
	slog.Info("Redirecting initial request to : " + location)
	c.Redirect(http.StatusFound, location)
}

//=============================================================================

func getRootSymbols(c *auth.Context) {
	code := c.GetCodeFromUrl()
	filter := c.Gin.Query("filter")
	if filter == "" {
		c.ReturnError(req.NewBadRequestError("No filter provided"))
	}

	res, err := business.GetRootSymbols(c, code, filter)
	if err == nil {
		_ = c.ReturnList(res, 0, 10000, len(res))
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getRootSymbol(c *auth.Context) {
	code := c.GetCodeFromUrl()
	root := c.Gin.Param("root")

	res, err := business.GetRootSymbol(c, code, root)
	if err == nil {
		_ = c.ReturnObject(res)
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getInstruments(c *auth.Context) {
	code := c.GetCodeFromUrl()
	root := c.Gin.Param("root")

	res, err := business.GetInstruments(c, code, root)
	if err == nil {
		_ = c.ReturnList(res, 0, 10000, len(res))
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getPriceBars(c *auth.Context) {
	code := c.GetCodeFromUrl()
	symbol := c.Gin.Param("symbol")
	date := c.Gin.Query("date")

	id, err := types.ParseIntDate(date, true)

	if err != nil {
		c.ReturnError(req.NewBadRequestError("Missing or invalid 'date' parameter"))
		return
	}

	res, err := business.GetPriceBars(c, code, symbol, id)
	if err == nil {
		_ = c.ReturnObject(res)
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getAccounts(c *auth.Context) {
	code := c.GetCodeFromUrl()

	res, err := business.GetAccounts(c, code)
	if err == nil {
		_ = c.ReturnList(res, 0, 1000, len(res))
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getOrders(c *auth.Context) {
	code := c.GetCodeFromUrl()

	res, err := business.GetOrders(c, code)
	if err == nil {
		_ = c.ReturnObject(res)
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func getPositions(c *auth.Context) {
	code := c.GetCodeFromUrl()

	res, err := business.GetPositions(c, code)
	if err == nil {
		_ = c.ReturnObject(res)
		return
	}

	c.ReturnError(err)
}

//=============================================================================

func testAdapter(c *auth.Context) {
	code := c.GetCodeFromUrl()
	tar := business.TestAdapterRequest{}
	err := c.BindParamsFromBody(&tar)

	if err == nil {
		var res string
		res, err = business.TestAdapter(c, code, &tar)
		if err == nil {
			_ = c.ReturnObject(res)
			return
		}
	}

	c.ReturnError(err)
}

//=============================================================================
