package scrapper

import (
	"sync"

	"github.com/junjie1999/exchange-cli/unit"
)

var URL = "www.x-rates.com"

type Currency struct {
	Timestamp     int64
	Source        string
	ExchangeRates map[string]map[string]float64
}

type currency struct {
	base      string
	timestamp int64
	Rates     map[string]float64
}

type Rate struct {
	mu      sync.Mutex
	RateMap map[string]*currency
}

func (Curr *Rate) add(base string, c *currency) {
	Curr.mu.Lock()
	defer Curr.mu.Unlock()
	Curr.RateMap[base] = c
}

func (Curr *Rate) toCurrencyMap() *Currency {

	Currency := &Currency{
		Timestamp:     Curr.RateMap["USD"].timestamp,
		Source:        URL,
		ExchangeRates: make(map[string]map[string]float64),
	}

	for base, curr := range Curr.RateMap {
		Currency.ExchangeRates[base] = curr.Rates
	}

	return Currency
}

func FetchData() *Currency {

	var wg sync.WaitGroup

	Curr := &Rate{
		RateMap: map[string]*currency{},
	}

	for _, base := range unit.X_RateUnits {
		wg.Add(1)
		go func(base string) {
			defer wg.Done()
			curr := scrapeURL(base)
			Curr.add(base, curr)
		}(base)
	}

	wg.Wait()

	return Curr.toCurrencyMap()
}
