//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package tradestation

//=============================================================================

const (
	UrlBrokerageAccounts  = "/v3/brokerage/accounts"
	UrlMarketDataSymbols  = "/v3/marketdata/symbols"
	UrlMarketDataBarcharts= "/v3/marketdata/barcharts"
	UrlSymbolsSearch      = "/v2/data/symbols/search"
	UrlSymbolsSuggest     = "/v2/data/symbols/suggest"
)

//=============================================================================
//=== Service: /v3/brokerage/accounts
//=============================================================================

type AccountsResponse struct {
	Accounts []Account
}

//=============================================================================

type Account struct {
	AccountID     string
	Currency      string
	Status        string
	AccountType   string
	AccountDetail AccountDetail
}

//=============================================================================

type AccountDetail struct {
	IsStockLocateEligible      bool
	EnrolledInRegTProgram      bool
	RequiresBuyingPowerWarning bool
	DayTradingQualified        bool
	OptionApprovalLevel        int
	PatternDayTrader           bool
}

//=============================================================================
//=== Service: /v3/brokerage/accounts/XXX/balances
//=============================================================================

type BalancesResponse struct {
	Balances []Balance
	Errors   []interface{}
}

//=============================================================================

type Balance struct {
	AccountID        string
	AccountType      string
	CashBalance      string
	BuyingPower      string
	Equity           string
	MarketValue      string
	TodaysProfitLoss string
	UnclearedDeposit string
	BalanceDetail    BalanceDetail
	CurrencyDetails  []CurrencyDetail
	Commission       string
}

//=============================================================================

type BalanceDetail struct {
	DayTradeExcess           string
	RealizedProfitLoss       string
	UnrealizedProfitLoss     string
	DayTradeOpenOrderMargin  string
	OpenOrderMargin          string
	DayTradeMargin           string
	InitialMargin            string
	MaintenanceMargin        string
	TradeEquity              string
	SecurityOnDeposit        string
	TodayRealTimeTradeEquity string
}

//=============================================================================

type CurrencyDetail struct {
	Currency              string
	Commission            string
	CashBalance           string
	RealizedProfitLoss    string
	UnrealizedProfitLoss  string
	InitialMargin         string
	MaintenanceMargin     string
	AccountConversionRate string
}

//=============================================================================
//=== Service: /v2/data/symbols/search/XXX
//=============================================================================

type SymbolFound struct {
	Name            string
	Description     string
	Exchange        string
	ExchangeID      int
	Category        string
	Country         string
	Root            string
	OptionType      string
	FutureType      string
	ExpirationDate  string
	ExpirationType  string
	StrikePrice     int
	Currency        string
	PointValue      int
	MinMove         float64
	DisplayType     int
	Error           interface{}
}

//=============================================================================
//=== Service: /v2/data/symbols/suggest/XXX
//=============================================================================

type RootFound struct {
	Country        string
	Currency       string
	Description    string
	Exchange       string
	ExpirationDate string
	LotSize        int
	MinMove        float64
	Name           string
	PointValue     float64
	Root           string
	StrikePrice    int
}

//=============================================================================
//=== Service: /v3/marketdata/symbols/XXX
//=============================================================================

type SymbolDetailsResponse struct {
	Symbols []SymbolDetails
	Errors  []interface{}
}

//=============================================================================

type SymbolDetails struct {
	AssetType      string
	Country        string
	Currency       string
	Description    string
	Exchange       string
	FutureType     string
	Symbol         string
	Root           string
	Underlying     string
	PriceFormat    PriceFormat
	QuantityFormat QuantityFormat
}

//=============================================================================

type PriceFormat struct {
	Format         string
	Decimals       string
	IncrementStyle string
	Increment      string
	PointValue     string
}

//=============================================================================

type QuantityFormat struct {
	Format               string
	Decimals             string
	IncrementStyle       string
	Increment            string
	MinimumTradeQuantity string
}

//=============================================================================
//=== Service: /v3/marketdata/barcharts/XXX
//=============================================================================

type BarchartsResponse struct {
	Bars    []Bar
	Error   string
	Message string
}

//=============================================================================

type Bar struct {
	TimeStamp    string
	Epoch        int64
	High         string
	Low          string
	Open         string
	Close        string
	UpVolume     int
	DownVolume   int
	UpTicks      int
	DownTicks    int
	OpenInterest string
}

//=============================================================================
