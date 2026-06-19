//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package adapter

import (
	"errors"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/algotiqa/types"
)

//=============================================================================
//=== Common connection parameters

const (
	ParamUsername  = "username"
	ParamPassword  = "password"
	ParamTwoFACode = "twoFACode"
)

//=============================================================================

type ParamType string

const (
	ParamTypeString   ParamType = "string"
	ParamTypePassword ParamType = "password"
	ParamTypeBool     ParamType = "bool"
	ParamTypeInt      ParamType = "int"
	ParamTypeList     ParamType = "list"
)

//=============================================================================

type ParamDef struct {
	Name     string    `json:"name"`
	Type     ParamType `json:"type"` // string|int|bool|password|list
	DefValue string    `json:"defValue"`
	Nullable bool      `json:"nullable"`
	MinValue int       `json:"minValue"`
	MaxValue int       `json:"maxValue"`
	Values   []string  `json:"values"`
}

//-----------------------------------------------------------------------------

func (p *ParamDef) Validate(values map[string]any) error {
	value, ok := values[p.Name]

	if !ok {
		//--- Check default value

		if p.DefValue != "" {
			switch p.Type {
			case ParamTypeBool:
				if p.DefValue != "true" && p.DefValue != "false" {
					return errors.New("invalid default value for a boolean parameter : " + p.Name)
				}
				break

			case ParamTypeInt:
				v, err := strconv.Atoi(p.DefValue)
				if err != nil {
					return errors.New("invalid value for an integer parameter : " + p.Name)
				}
				if v < p.MinValue || v > p.MaxValue {
					return errors.New("invalid range for this integer parameter : " + p.Name)
				}
				break
			}
			return nil
		}

		if !p.Nullable {
			return errors.New("missing mandatory value for parameter : " + p.Name)
		}
	} else {
		//--- Check provided value

		t := reflect.TypeOf(value)
		slog.Info("Param Type is", "type", t.Name())
		switch t.Name() {
		case "string":
			if p.Type == ParamTypeString || p.Type == ParamTypePassword || p.Type == ParamTypeList {
				return nil
			}
			break

		case "bool":
			if p.Type == ParamTypeBool {
				return nil
			}
			break

		case "int":
			if p.Type == ParamTypeInt {
				return nil
			}

		default:
			return errors.New("unknown parameter type : " + p.Name)
		}

		return errors.New("invalid parameter value : " + p.Name)
	}

	return nil
}

//=============================================================================

type Info struct {
	Code                 string      `json:"code"`
	Name                 string      `json:"name"`
	SupportsData         bool        `json:"supportsData"`
	SupportsBroker       bool        `json:"supportsBroker"`
	SupportsMultipleData bool        `json:"supportsMultipleData"`
	SupportsInventory    bool        `json:"supportsInventory"`
	ConfigParams         []*ParamDef `json:"configParams"`
	ConnectParams        []*ParamDef `json:"connectParams"`
}

//=============================================================================

type ConnectionResult struct {
	Status  ContextStatus `json:"status"`
	Message string        `json:"message"`
	Url     string        `json:"url"`
	Params  []*ParamDef   `json:"params"`
}

//=============================================================================

type Adapter interface {
	GetInfo() *Info
	GetAuthUrl() string
	Clone(configParams map[string]any, connectParams map[string]any) Adapter
	GetConnectParams(configParams map[string]any) []*ParamDef
	Connect(ctx *ConnectionContext) *ConnectionResult
	Disconnect(ctx *ConnectionContext) error
	IsWebLoginCompleted(httpCode int, path string) bool
	InitFromWebLogin(reqHeader *http.Header, resCookies []*http.Cookie) error
	GetTokenExpSeconds() int
	RefreshToken() error

	//--- Services

	GetRootSymbols(filter string) ([]*RootSymbol, error)
	GetRootSymbol(root string) (*RootSymbol, error)
	GetInstruments(root string) ([]*Instrument, error)
	GetPriceBars(symbol string, date types.Date) (*PriceBars, error)
	GetAccounts() ([]*Account, error)
	GetOrders() (any, error)
	GetPositions() (any, error)
	TestService(path, param string) (string, error)
}

//=============================================================================
//===
//=== API model
//===
//=============================================================================

type AccountType int

const (
	AccountTypeFutures = 0
	AccountTypeCrypto  = 1
)

//-----------------------------------------------------------------------------

type Account struct {
	Code                 string      `json:"code"`
	Type                 AccountType `json:"type"`
	CurrencyCode         string      `json:"currencyCode"`
	CashBalance          float64     `json:"cashBalance"`
	Equity               float64     `json:"equity"`
	RealizedProfitLoss   float64     `json:"realizedProfitLoss"`
	UnrealizedProfitLoss float64     `json:"unrealizedProfitLoss"`
	OpenOrderMargin      float64     `json:"openOrderMargin"`
	InitialMargin        float64     `json:"initialMargin"`
	MaintenanceMargin    float64     `json:"maintenanceMargin"`
}

//=============================================================================

type Instrument struct {
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Exchange       string     `json:"exchange"`
	Country        string     `json:"country"`
	Root           string     `json:"root"`
	ExpirationDate *time.Time `json:"expirationDate"`
	PointValue     int        `json:"pointValue"`
	MinMove        float64    `json:"minMove"`
	Continuous     bool       `json:"continuous"`
	Month          string     `json:"month"`
}

//=============================================================================

type RootSymbol struct {
	Code       string  `json:"code"`
	Instrument string  `json:"instrument"`
	Exchange   string  `json:"exchange"`
	PointValue float64 `json:"pointValue"`
	Increment  float64 `json:"increment"`
	Country    string  `json:"country"`
	Currency   string  `json:"currency"`
}

//=============================================================================

type PriceBars struct {
	Symbol          string      `json:"symbol"`
	Date            int         `json:"date"`
	Days            int         `json:"days"`
	Bars            []*PriceBar `json:"bars"`
	NoData          bool        `json:"noData"`
	Timeout         bool        `json:"timeout"`
	TooManyRequests bool        `json:"tooManyRequests"`
}

//=============================================================================

type PriceBar struct {
	TimeStamp    time.Time
	High         float64
	Low          float64
	Open         float64
	Close        float64
	UpVolume     int
	DownVolume   int
	UpTicks      int
	DownTicks    int
	OpenInterest int
}

//=============================================================================
