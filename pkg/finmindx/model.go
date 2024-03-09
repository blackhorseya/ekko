package finmindx

import (
	"encoding/json"
	"time"

	"github.com/blackhorseya/ekko/pkg/timex"
)

// Response is used to represent the response.
type Response struct {
	Message string      `json:"msg"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// TaiwanStockPrice is used to represent the Taiwan stock price.
type TaiwanStockPrice struct {
	Date            time.Time `json:"date"`
	StockID         string    `json:"stock_id"`
	TradingVolume   int64     `json:"Trading_Volume"`
	TradingMoney    int64     `json:"Trading_money"`
	Open            float64   `json:"open"`
	Max             float64   `json:"max"`
	Min             float64   `json:"min"`
	Close           float64   `json:"close"`
	Spread          float64   `json:"spread"`
	TradingTurnover float64   `json:"Trading_turnover"`
}

func (x *TaiwanStockPrice) UnmarshalJSON(bytes []byte) error {
	type Alias TaiwanStockPrice
	aux := &struct {
		*Alias
		Date string `json:"date,omitempty"`
	}{
		Alias: (*Alias)(x),
	}

	err := json.Unmarshal(bytes, &aux)
	if err != nil {
		return err
	}

	var date time.Time
	if len(aux.Date) != 0 && aux.Date != "None" {
		date, err = time.ParseInLocation(dateFormat, aux.Date, timex.LocTaipei)
		if err != nil {
			return err
		}

		x.Date = date
	}

	return nil
}

// TaiwanStockPriceResponse is used to represent the Taiwan stock price response.
type TaiwanStockPriceResponse struct {
	*Response `json:",inline"`
	Data      []*TaiwanStockPrice `json:"data,omitempty"`
}

// TaiwanStockInfo is used to represent the Taiwan stock info.
type TaiwanStockInfo struct {
	IndustryCategory string    `json:"industry_category"`
	StockID          string    `json:"stock_id"`
	StockName        string    `json:"stock_name"`
	Type             string    `json:"type"`
	Date             time.Time `json:"date"`
}

func (x *TaiwanStockInfo) UnmarshalJSON(bytes []byte) error {
	type Alias TaiwanStockInfo
	aux := &struct {
		*Alias
		Date string `json:"date,omitempty"`
	}{
		Alias: (*Alias)(x),
	}

	err := json.Unmarshal(bytes, &aux)
	if err != nil {
		return err
	}

	var date time.Time
	if len(aux.Date) != 0 && aux.Date != "None" {
		date, err = time.ParseInLocation(dateFormat, aux.Date, timex.LocTaipei)
		if err != nil {
			return err
		}

		x.Date = date
	}

	return nil
}
