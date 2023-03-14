package setbk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

const (
	baseURL       = "https://www.settrade.com/api/set"
	defaultMarket = "set"

	floatLayout = "+#,###.##"
	timeLayout  = "Jan 2 2006"
)

func GetPrettySummary(market string) PrettySummary {
	s := GetSummary(market)

	return PrettySummary{
		Index:           prettyIndex(s.Index),
		InvestorSummary: prettyInvestorSummary(s.InvestorSummary),
		Rankings:        prettyRankings(s.Rankings),
	}
}

func prettyIndex(i index) pIndex {
	dt, err := time.Parse(time.RFC3339, i.MarketDateTime)
	if err != nil {
		log.Println(err)
		return pIndex{}
	}
	i.MarketDateTime = dt.Format(timeLayout)

	return pIndex{
		Symbol:         strings.ToUpper(i.Symbol),
		NameEN:         strings.ToUpper(i.NameEN),
		NameTH:         strings.ToUpper(i.NameTH),
		Prior:          humanize.FormatFloat("", i.Prior),
		High:           humanize.FormatFloat("", i.High),
		Low:            humanize.FormatFloat("", i.Low),
		Last:           humanize.FormatFloat("", i.Last),
		Change:         humanize.FormatFloat(floatLayout, i.Change),
		PercentChange:  humanize.FormatFloat(floatLayout, i.PercentChange),
		Volume:         humanize.FormatFloat("", i.Volume/1_000_000),
		Value:          humanize.FormatFloat("", i.Value/1_000_000),
		QuerySymbol:    i.QuerySymbol,
		MarketStatus:   i.MarketStatus,
		MarketDateTime: i.MarketDateTime,
		MarketName:     i.MarketName,
		IndustryName:   i.IndustryName,
		SectorName:     i.SectorName,
		Level:          i.Level,
	}
}

func prettyInvestorSummary(is investorSummary) pInvestorSummary {
	var investors []pInvestor
	for _, inv := range is.Investors {
		investors = append(investors, pInvestor{
			Type:             strings.ToUpper(inv.Type),
			BuyValue:         humanize.FormatFloat("", inv.BuyValue),
			SellValue:        humanize.FormatFloat("", inv.SellValue),
			NetValue:         humanize.FormatFloat(floatLayout, inv.NetValue/1_000_000),
			PercentBuyValue:  humanize.FormatFloat(floatLayout, inv.PercentBuyValue),
			PercentSellValue: humanize.FormatFloat(floatLayout, inv.PercentSellValue),
		})
	}

	return pInvestorSummary{
		Name:       is.Name,
		AsOfDate:   is.AsOfDate,
		BeginDate:  is.BeginDate,
		EndDate:    is.EndDate,
		TotalValue: humanize.FormatFloat("", is.TotalValue),
		Investors:  investors,
	}
}

func prettyRankings(rr []ranking) []pRanking {
	var rankings []pRanking
	for _, r := range rr {
		rankings = append(rankings, pRanking{
			RankingType:    r.RankingType,
			Market:         r.Market,
			SecurityType:   r.SecurityType,
			MarketDateTime: r.MarketDateTime,
			RankingPeriod:  r.RankingPeriod,
			Stocks:         prettyStocks(r.Stocks),
		})
	}

	return rankings
}

func prettyStocks(ss []stock) []pStock {
	var stocks []pStock
	for _, s := range ss {
		stocks = append(stocks, pStock{
			Symbol:        s.Symbol,
			Sign:          s.Sign,
			Prior:         humanize.FormatFloat("", s.Prior),
			Last:          humanize.FormatFloat("", s.Last),
			Change:        humanize.FormatFloat(floatLayout, s.Change),
			PercentChange: humanize.FormatFloat(floatLayout, s.PercentChange),
			High:          humanize.FormatFloat("", s.High),
			Low:           humanize.FormatFloat("", s.Low),
			TotalVolume:   humanize.FormatFloat("", s.TotalVolume),
			TotalValue:    humanize.FormatFloat("", s.TotalValue),
			AomVolume:     humanize.FormatFloat("", s.AomVolume),
			AomValue:      humanize.FormatFloat("", s.AomValue),
			RankingValue:  humanize.FormatFloat("", s.RankingValue),
		})
	}

	return stocks
}

var rankingTypes = []string{"mostActiveValue", "topGainer", "topLoser"}

func GetSummary(market string) Summary {
	var rankings []ranking
	for _, rt := range rankingTypes {
		rankings = append(rankings, getRanking(market, rt))
	}

	return Summary{
		Index:           getIndex(market),
		InvestorSummary: getInvestorSummary(market),
		Rankings:        rankings,
	}
}

func getIndex(market string) index {
	resp, err := http.Get(indexURL(market))
	if err != nil {
		log.Println(err)
		return index{}
	}
	if resp.StatusCode != http.StatusOK {
		return index{}
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return index{}
	}

	var res index
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return index{}
	}

	return res
}

func getInvestorSummary(market string) investorSummary {
	resp, err := http.Get(investorSummaryURL(market))
	if err != nil {
		log.Println(err)
		return investorSummary{}
	}
	if resp.StatusCode != http.StatusOK {
		return investorSummary{}
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return investorSummary{}
	}

	var res investorSummary
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return investorSummary{}
	}

	return res
}

func getRanking(market, rType string) ranking {
	resp, err := http.Get(rankingURL(market, rType))
	if err != nil {
		log.Println(err)
		return ranking{}
	}
	if resp.StatusCode != http.StatusOK {
		return ranking{}
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return ranking{}
	}

	var res ranking
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return ranking{}
	}

	return res
}

func indexURL(market string) string {
	if market == "" {
		market = defaultMarket
	}
	return fmt.Sprintf("%s/index/%s/info", baseURL, market)
}

func investorSummaryURL(market string) string {
	if market == "" {
		market = defaultMarket
	}
	return fmt.Sprintf("%s/market/%s/investor-type", baseURL, market)
}

func rankingURL(market, rType string) string {
	if market == "" {
		market = defaultMarket
	}
	return fmt.Sprintf("%s/ranking/%s/%s/s?count=5", baseURL, rType, market)
}
