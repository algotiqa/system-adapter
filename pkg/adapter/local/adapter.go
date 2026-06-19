//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package local

import (
	"net/http"

	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/types"
)

//=============================================================================

var configParams []*adapter.ParamDef

//-----------------------------------------------------------------------------

var info = adapter.Info{
	Code:                 "LOCAL",
	Name:                 "Local system",
	ConfigParams:         configParams,
	SupportsData:         true,
	SupportsBroker:       true,
	SupportsMultipleData: true,
	SupportsInventory:    false,
}

//=============================================================================

func NewAdapter() adapter.Adapter {
	return &local{}
}

//=============================================================================

type local struct {
}

//=============================================================================

func (a *local) GetInfo() *adapter.Info {
	return &info
}

//=============================================================================

func (a *local) GetAuthUrl() string {
	return ""
}

//=============================================================================

func (a *local) Clone(configParams map[string]any, connectParams map[string]any) adapter.Adapter {
	b := *a
	return &b
}

//=============================================================================

func (a *local) GetConnectParams(configParams map[string]any) []*adapter.ParamDef {
	return []*adapter.ParamDef{}
}

//=============================================================================

func (a *local) Connect(ctx *adapter.ConnectionContext) *adapter.ConnectionResult {
	return &adapter.ConnectionResult{
		Status: adapter.ContextStatusConnected,
	}
}

//=============================================================================

func (a *local) Disconnect(ctx *adapter.ConnectionContext) error {
	return nil
}

//=============================================================================

func (a *local) IsWebLoginCompleted(httpCode int, path string) bool {
	return true
}

//=============================================================================

func (a *local) InitFromWebLogin(reqHeader *http.Header, resCookies []*http.Cookie) error {
	return nil
}

//=============================================================================

func (a *local) GetTokenExpSeconds() int {
	return 0
}

//=============================================================================

func (a *local) RefreshToken() error {
	return nil
}

//=============================================================================
//===
//=== Services
//===
//=============================================================================

func (a *local) GetRootSymbols(filter string) ([]*adapter.RootSymbol, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetRootSymbol(root string) (*adapter.RootSymbol, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetInstruments(root string) ([]*adapter.Instrument, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetPriceBars(symbol string, date types.Date) (*adapter.PriceBars, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetAccounts() ([]*adapter.Account, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetOrders() (any, error) {
	return nil, nil
}

//=============================================================================

func (a *local) GetPositions() (any, error) {
	return nil, nil
}

//=============================================================================

func (a *local) TestService(path, param string) (string, error) {
	return "", nil
}

//=============================================================================
