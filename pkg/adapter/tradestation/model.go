//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

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
	"net/http"

	"github.com/algotiqa/system-adapter/pkg/adapter"
)

//=============================================================================
//--- Config parameters

const (
	//--- Config params

	ParamAccount  = "account"
	ParamAuthType = "authType"

	//--- Connection params (authType="browser" )

	ParamClientId     = "clientId"
	ParamClientSecret = "clientSecret"
	ParamClientCode   = "clientCode"

	//--- List options

	ParamAccountTest = "test"
	ParamAccountLive = "live"

	ParamAuthTypeBrowser  = "browser"
	ParamAuthTypeInternal = "internal"
)

//=============================================================================

const (
	LiveAPI = "https://api.tradestation.com"
	DemoAPI = "https://sim-api.tradestation.com"

	AuthorizeUrl  = "https://signin.tradestation.com/authorize"
	OauthTokenUrl = "https://signin.tradestation.com/oauth/token"

	LoginPageUrl       = "https://my.tradestation.com/api/auth/login?returnTo=%2F"
	LoginPostUrl       = "https://signin.tradestation.com/usernamepassword/login"
	LoginCallbackUrl   = "https://signin.tradestation.com/login/callback"
	LoginTwoFAUrl      = "https://signin.tradestation.com/u/mfa-otp-challenge"
	LoginTwoFAPath     = "/u/mfa-otp-challenge"
	LoginDashboardPath = "/dashboard"
	RefreshTokenUrl    = "https://my.tradestation.com/api/auth/token"
)

//=============================================================================

var configParams = []*adapter.ParamDef{
	{
		Name:     ParamAccount,
		Type:     adapter.ParamTypeList,
		DefValue: ParamAccountTest,
		Nullable: false,
		Values:   []string{ParamAccountTest, ParamAccountLive},
	},
	{
		Name:     ParamAuthType,
		Type:     adapter.ParamTypeList,
		DefValue: ParamAuthTypeBrowser,
		Nullable: false,
		Values:   []string{ParamAuthTypeBrowser, ParamAuthTypeInternal},
	},
}

//-----------------------------------------------------------------------------

var connectParamsBrowser = []*adapter.ParamDef{
	{
		Name:     ParamClientId,
		Type:     adapter.ParamTypeString,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 64,
	},
	{
		Name:     ParamClientSecret,
		Type:     adapter.ParamTypeString,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 64,
	},
}

//-----------------------------------------------------------------------------

var connectParamsInternal = []*adapter.ParamDef{
	{
		Name:     adapter.ParamUsername,
		Type:     adapter.ParamTypeString,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 64,
	},
	{
		Name:     adapter.ParamPassword,
		Type:     adapter.ParamTypePassword,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 64,
	},
	{
		Name:     adapter.ParamTwoFACode,
		Type:     adapter.ParamTypeString,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 64,
	},
}

//-----------------------------------------------------------------------------

var connectParamsCode = []*adapter.ParamDef{
	{
		Name:     ParamClientCode,
		Type:     adapter.ParamTypeString,
		DefValue: "",
		Nullable: false,
		MinValue: 0,
		MaxValue: 128,
	},
}

//-----------------------------------------------------------------------------

var info = adapter.Info{
	Code:                 "TS",
	Name:                 "Tradestation",
	ConfigParams:         configParams,
	SupportsData:         true,
	SupportsBroker:       true,
	SupportsMultipleData: false,
	SupportsInventory:    true,
}

//=============================================================================

type ConfigParams struct {
	Account  string
	AuthType string
}

//=============================================================================

type ConnectParams struct {
	Username     string
	Password     string
	TwoFACode    string
	ClientId     string
	ClientSecret string
	ClientCode   string
}

//=============================================================================

type tradestation struct {
	configParams  *ConfigParams
	connectParams *ConnectParams
	client        *http.Client
	header        *http.Header
	accessToken   string
	refreshToken  string
	clientId      string
	apiUrl        string
}

//=============================================================================
