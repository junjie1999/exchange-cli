package scrapper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"

	"github.com/junjie1999/exchange-cli/unit"
)

type scrapper struct {
	collect *colly.Collector
}

func NewScrapper() *scrapper {
	c := colly.NewCollector(
		colly.AllowedDomains(URL),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	return &scrapper{
		collect: c,
	}
}

func (sp *scrapper) setupCollector(Curr *currency) {
	sp.setupOnHTML(Curr)
	sp.setupOnError()
}

func (sp *scrapper) setupOnHTML(curr *currency) {
	sp.collect.OnHTML(".moduleContent", func(e *colly.HTMLElement) {
		sp.parseCurrency(e, curr)
	})
}

func (sp *scrapper) setupOnError() {
	sp.collect.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == 403 {
			fmt.Errorf("access forbidden (403). Consider increasing delays or using proxies")
		} else if r.StatusCode == 404 {
			fmt.Errorf("page not found (404)")
		} else {
			fmt.Errorf(err.Error())
		}
	})
}

func (sp *scrapper) parseCurrency(e *colly.HTMLElement, Curr *currency) {

	e.ForEachWithBreak(".ratesTimestamp", func(_ int, eTime *colly.HTMLElement) bool {
		_, Curr.timestamp = convertToTimestamp(
			strings.TrimSpace(eTime.Text))
		return false
	})

	e.ForEach(".tablesorter.ratesTable>tbody>tr", func(_ int, row *colly.HTMLElement) {

		var target string

		row.ForEach("td", func(i int, col *colly.HTMLElement) {

			switch i {
			case 0:
				target = unit.GetXRateUnit(
					strings.TrimSpace(col.Text))
			case 1:
				if rate, err := strconv.ParseFloat(col.Text, 64); err == nil {
					Curr.Rates[target] = rate
				} else {
					Curr.Rates[target] = -1
				}
			}
		}) //row.ForEach "td"
	})
}

func (sp *scrapper) visitURL(base string) {
	url := fmt.Sprintf("https://www.x-rates.com/table/?from=%s&amount=1", base)
	sp.collect.Visit(url)
}

func scrapeURL(base string) *currency {

	Curr := &currency{
		base:  base,
		Rates: make(map[string]float64),
	}

	scraper := NewScrapper()
	scraper.setupCollector(Curr) // Add this line to set up the collector

	scraper.visitURL(base)

	return Curr
}

func convertToTimestamp(xRateTime string) (error, int64) {
	layout := "Jan 02, 2006 15:04 MST"
	t, err := time.Parse(layout, xRateTime)
	if err != nil {
		return err, 0
	}
	return nil, t.Unix()
}
