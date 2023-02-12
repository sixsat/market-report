package set

type Summary struct {
	Index           index
	InvestorSummary investorSummary
	Rankings        []ranking
}

type index struct {
	Symbol         string  `json:"symbol"`
	NameEN         string  `json:"nameEn"`
	NameTH         string  `json:"nameTh"`
	Prior          float64 `json:"prior"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Last           float64 `json:"last"`
	Change         float64 `json:"change"`
	PercentChange  float64 `json:"percentChange"`
	Volume         float64 `json:"volume"`
	Value          float64 `json:"value"`
	QuerySymbol    string  `json:"querySymbol"`
	MarketStatus   string  `json:"marketStatus"`
	MarketDateTime string  `json:"marketDateTime"`
	MarketName     string  `json:"marketName"`
	IndustryName   string  `json:"industryName"`
	SectorName     string  `json:"sectorName"`
	Level          string  `json:"level"`
}

type investorSummary struct {
	Name       string     `json:"name"`
	AsOfDate   string     `json:"asOfDate"`
	BeginDate  string     `json:"beginDate"`
	EndDate    string     `json:"endDate"`
	TotalValue float64    `json:"totalValue"`
	Investors  []investor `json:"investors"`
}

type investor struct {
	Type             string  `json:"type"`
	BuyValue         float64 `json:"buyValue"`
	SellValue        float64 `json:"sellValue"`
	NetValue         float64 `json:"netValue"`
	PercentBuyValue  float64 `json:"percentBuyValue"`
	PercentSellValue float64 `json:"percentSellValue"`
}

type ranking struct {
	RankingType    string      `json:"rankingType"`
	Market         string      `json:"market"`
	SecurityType   string      `json:"securityType"`
	MarketDateTime string      `json:"marketDateTime"`
	RankingPeriod  interface{} `json:"rankingPeriod"`
	Stocks         []stock     `json:"stocks"`
}

type stock struct {
	Symbol        string  `json:"symbol"`
	Sign          string  `json:"sign"`
	Prior         float64 `json:"prior"`
	Last          float64 `json:"last"`
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percentChange"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	TotalVolume   float64 `json:"totalVolume"`
	TotalValue    float64 `json:"totalValue"`
	AomVolume     float64 `json:"aomVolume"`
	AomValue      float64 `json:"aomValue"`
	RankingValue  float64 `json:"rankingValue"`
}
