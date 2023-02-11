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

func GetIndex(market string) *Index {
	resp, err := http.Get(indexURL(market))
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var res Index
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &res
}

func GetSummary(market string) *Summary {
	resp, err := http.Get(summaryURL(market))
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var res Summary
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &res
}

func GetRanking(market, rType string) *Ranking {
	resp, err := http.Get(rankingURL(market, rType))
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var res Ranking
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &res
}

func indexURL(market string) string {
	if market == "" {
		market = defaultMarket
	}
	return fmt.Sprintf("%s/index/%s/info", baseURL, market)
}

func summaryURL(market string) string {
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
