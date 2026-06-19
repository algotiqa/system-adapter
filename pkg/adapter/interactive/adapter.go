//=============================================================================
//===
//=== Copyright (C) 2024-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package interactive

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/types"
)

//=============================================================================

func NewAdapter() adapter.Adapter {
	return &ib{}
}

//=============================================================================

func (a *ib) GetInfo() *adapter.Info {
	return &info
}

//=============================================================================

func (a *ib) GetAuthUrl() string {
	return a.configParams.AuthUrl
}

//=============================================================================

func (a *ib) Clone(configParams map[string]any, connectParams map[string]any) adapter.Adapter {
	b := *a
	b.configParams = retrieveParams(configParams)
	return &b
}

//=============================================================================

func (a *ib) GetConnectParams(configParams map[string]any) []*adapter.ParamDef {
	return []*adapter.ParamDef{}
}

//=============================================================================

func (a *ib) Connect(ctx *adapter.ConnectionContext) *adapter.ConnectionResult {
	if a.configParams.NoAuth {
		//TODO: we should check if the connection actually works...
		//---   connection to the gateway

		return &adapter.ConnectionResult{
			Status: adapter.ContextStatusConnected,
		}
	}

	return &adapter.ConnectionResult{
		//--- Proxy url here
	}
}

//=============================================================================

func (a *ib) Disconnect(ctx *adapter.ConnectionContext) error {
	return nil
}

//=============================================================================

func (a *ib) IsWebLoginCompleted(httpCode int, path string) bool {
	return httpCode == http.StatusFound && path == "/sso/Dispatcher"
}

//=============================================================================

func (a *ib) InitFromWebLogin(reqHeader *http.Header, resCookies []*http.Cookie) error {
	header, err := buildHttpHeader(reqHeader, resCookies)
	if err != nil {
		return err
	}

	a.header = header
	a.client = &http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	res, err := a.ssoValidate()

	if err != nil {
		return err
	}

	if !res.Result {
		return errors.New("session is invalid")
	}

	//	orders,err := a.getAccountOrders()
	//	pnl,err := a.getAccountProfitAndLoss()
	//	tic,err := a.tickle()
	//	fmt.Println("RES:"+tic.Session)
	return nil
}

//=============================================================================

func (a *ib) GetTokenExpSeconds() int {
	return 0
}

//=============================================================================

func (a *ib) RefreshToken() error {
	return nil
}

//=============================================================================
//===
//=== Services
//===
//=============================================================================

func (a *ib) GetRootSymbols(filter string) ([]*adapter.RootSymbol, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetRootSymbol(root string) (*adapter.RootSymbol, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetInstruments(root string) ([]*adapter.Instrument, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetPriceBars(symbol string, date types.Date) (*adapter.PriceBars, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetAccounts() ([]*adapter.Account, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetOrders() (any, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) GetPositions() (any, error) {
	return nil, nil
}

//=============================================================================

func (a *ib) TestService(path, param string) (string, error) {
	return "", nil
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func retrieveParams(values map[string]any) *Params {
	return &Params{
		AuthUrl: values[ParamAuthUrl].(string),
		ApiUrl:  values[ParamApiUrl].(string),
	}
}

//=============================================================================

func buildHttpHeader(reqHeader *http.Header, resCookies []*http.Cookie) (*http.Header, error) {
	userId, err := findUserId(resCookies)
	if err != nil {
		return nil, err
	}

	cookies := reqHeader.Get("Cookie")
	cookies += "; " + userId.Name + "=" + userId.Value

	h := make(http.Header)
	h.Set("Accept", "*/*")
	h.Set("Cache-Control", "no-cache")
	h.Set("Pragma", "no-cache")
	h.Set("Sec-Fetch-Dest", "empty")
	h.Set("Sec-Fetch-Mode", "cors")

	h.Set("Cookie", cookies)
	h.Set("Accept-Encoding", reqHeader.Get("Accept-Encoding"))
	h.Set("Accept-Language", reqHeader.Get("Accept-Language"))

	return &h, nil
}

//=============================================================================

func findUserId(cookies []*http.Cookie) (*http.Cookie, error) {
	for _, cookie := range cookies {
		if cookie.Name == "USERID" {
			return cookie, nil
		}
	}

	return nil, errors.New("cookie USERID was not found")
}

//=============================================================================

func (a *ib) doGet(url string, output any) error {
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Error creating a GET request", "error", err.Error())
		return err
	}

	rq.Header = *a.header
	res, err := a.client.Do(rq)
	return req.BuildResponse(res, err, &output)
}

//=============================================================================

func (a *ib) doPost(url string, params any, output any) error {
	body, err := json.Marshal(&params)
	if err != nil {
		slog.Error("Error marshalling POST parameter", "error", err.Error())
		return err
	}

	reader := bytes.NewReader(body)

	rq, err := http.NewRequest("POST", url, reader)
	if err != nil {
		slog.Error("Error creating a POST request", "error", err.Error())
		return err
	}
	rq.Header = *a.header
	rq.Header.Set("Content-Type", "Application/json")

	res, err := a.client.Do(rq)
	return req.BuildResponse(res, err, &output)
}

//=============================================================================
//===
//=== IBKR services
//===
//=============================================================================

func (a *ib) ssoValidate() (*Validate, error) {
	apiUrl := a.configParams.ApiUrl + "/v1/api/sso/validate"
	var res Validate
	err := a.doGet(apiUrl, &res)

	return &res, err
}

//=============================================================================

func (a *ib) getAccountOrders() (*OrdersResponse, error) {
	apiUrl := a.configParams.ApiUrl + "/v1/api/iserver/account/orders?force=true"
	var res OrdersResponse
	err := a.doGet(apiUrl, &res)

	return &res, err
}

//=============================================================================

func (a *ib) getAccountProfitAndLoss() (*AccountPnLResponse, error) {
	apiUrl := a.configParams.ApiUrl + "/v1/api/iserver/account/pnl/partitioned"
	var res AccountPnLResponse
	err := a.doGet(apiUrl, &res)

	return &res, err
}

//=============================================================================

func (a *ib) tickle() (*TickleResponse, error) {
	apiUrl := a.configParams.ApiUrl + "/v1/api/tickle"
	var res TickleResponse
	err := a.doPost(apiUrl, "{}", &res)

	return &res, err
}

//=============================================================================
