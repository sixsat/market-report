package set

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://www.settrade.com/api/set"

	defaultMarket = "set"
)

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
