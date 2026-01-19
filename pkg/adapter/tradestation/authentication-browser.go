//=============================================================================
/*
Copyright © 2026 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package tradestation

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/algotiqa/core/req"
	"github.com/algotiqa/system-adapter/pkg/adapter"
)

//=============================================================================
//===
//=== Model
//===
//=============================================================================

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

//=============================================================================
//===
//=== Methods
//===
//=============================================================================

func (a *tradestation) connectBrowser(ctx *adapter.ConnectionContext) *adapter.ConnectionResult {
	if a.connectParams.ClientCode == "" {
		return &adapter.ConnectionResult{
			Status: adapter.ContextStatusConnecting,
			Params: connectParamsCode,
			Url:    buildUrl(a.connectParams.ClientId),
		}
	}

	tr, err := a.getNewTokens()
	if err != nil {
		return connectError(err)
	}
	if tr.RefreshToken == "" {
		err = errors.New("empty refresh token")
		return connectError(err)
	}

	a.accessToken = tr.AccessToken
	a.refreshToken = tr.RefreshToken

	a.setupAPIUrl()

	//--- Test tokens & accounts
	err = a.testToken()
	if err != nil {
		return connectError(err)
	}

	return &adapter.ConnectionResult{
		Status: adapter.ContextStatusConnected,
	}
}

//=============================================================================

func (a *tradestation) refreshTokenBrowser() error {
	var params = url.Values{}
	params.Set("grant_type", "refresh_token")
	params.Set("client_id", a.connectParams.ClientId)
	params.Set("client_secret", a.connectParams.ClientSecret)
	params.Set("refresh_token", a.refreshToken)

	tr, err := a.getTokens(params)
	if err == nil {
		a.accessToken = tr.AccessToken

		if a.accessToken == "" {
			err = errors.New("empty access token (refresh token is not working)")
		}
	}
	return err
}

//=============================================================================

func buildUrl(clientId string) string {
	var sb strings.Builder
	sb.WriteString(AuthorizeUrl)
	sb.WriteString("?response_type=code")
	sb.WriteString("&client_id=" + clientId)
	sb.WriteString("&redirect_uri=http%3A%2F%2Flocalhost%3A8080")
	sb.WriteString("&audience=https%3A%2F%2Fapi.tradestation.com")
	sb.WriteString("&state=STATE")
	sb.WriteString("&scope=openid%20offline_access%20profile%20MarketData%20ReadAccount%20Trade%20Matrix")

	return sb.String()
}

//=============================================================================

func (a *tradestation) getNewTokens() (*TokenResponse, error) {
	var params = url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("client_id", a.connectParams.ClientId)
	params.Set("client_secret", a.connectParams.ClientSecret)
	params.Set("code", a.connectParams.ClientCode)
	params.Set("redirect_uri", "http://localhost:8080")

	return a.getTokens(params)
}

//=============================================================================

func (a *tradestation) getTokens(params url.Values) (*TokenResponse, error) {
	payload := bytes.NewBufferString(params.Encode())

	rq, err := http.NewRequest("POST", OauthTokenUrl, payload)
	if err != nil {
		slog.Error("getTokens: Error creating a POST request", "error", err.Error())
		return nil, err
	}

	rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	out := &TokenResponse{}
	res, err := a.client.Do(rq)
	err = req.BuildResponse(res, err, out)

	if res != nil && res.Body != nil {
		_ = res.Body.Close()
	}

	return out, err
}

//=============================================================================
