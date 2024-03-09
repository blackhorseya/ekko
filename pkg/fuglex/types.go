package fuglex

type IntradayQuote struct {
	Date           string    `json:"date"`
	Type           string    `json:"type"`
	Exchange       string    `json:"exchange"`
	Market         string    `json:"market"`
	Symbol         string    `json:"symbol"`
	Name           string    `json:"name"`
	ReferencePrice float64   `json:"referencePrice"`
	PreviousClose  float64   `json:"previousClose"`
	OpenPrice      float64   `json:"openPrice"`
	OpenTime       int64     `json:"openTime"`
	HighPrice      float64   `json:"highPrice"`
	HighTime       int64     `json:"highTime"`
	LowPrice       float64   `json:"lowPrice"`
	LowTime        int64     `json:"lowTime"`
	ClosePrice     float64   `json:"closePrice"`
	CloseTime      int64     `json:"closeTime"`
	AvgPrice       float64   `json:"avgPrice"`
	Change         float64   `json:"change"`
	ChangePercent  float64   `json:"changePercent"`
	Amplitude      float64   `json:"amplitude"`
	LastPrice      float64   `json:"lastPrice"`
	LastSize       float64   `json:"lastSize"`
	Bids           []Bids    `json:"bids"`
	Asks           []Asks    `json:"asks"`
	Total          Total     `json:"total"`
	LastTrade      LastTrade `json:"lastTrade"`
	LastTrial      LastTrial `json:"lastTrial"`
	IsClose        bool      `json:"isClose"`
	Serial         int       `json:"serial"`
	LastUpdated    int64     `json:"lastUpdated"`
}

type Bids struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Asks struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Total struct {
	TradeValue       int64 `json:"tradeValue"`
	TradeVolume      int   `json:"tradeVolume"`
	TradeVolumeAtBid int   `json:"tradeVolumeAtBid"`
	TradeVolumeAtAsk int   `json:"tradeVolumeAtAsk"`
	Transaction      int   `json:"transaction"`
	Time             int64 `json:"time"`
}

type LastTrade struct {
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	Price  float64 `json:"price"`
	Size   float64 `json:"size"`
	Time   int64   `json:"time"`
	Serial int     `json:"serial"`
}

type LastTrial struct {
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	Price  float64 `json:"price"`
	Size   float64 `json:"size"`
	Time   int64   `json:"time"`
	Serial int     `json:"serial"`
}
