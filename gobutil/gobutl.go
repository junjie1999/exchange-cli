package gobutil

import (
	"encoding/gob"
	"os"

	"github.com/junjie1999/exchange-cli/scrapper"
)

type ExchangeRates *scrapper.Currency

var CurrencyFile = "currency.gob"

func UpdateCurrency() (error, *int64) {

	CMap := scrapper.FetchData()

	file, err := os.Create(CurrencyFile)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	enc := gob.NewEncoder(file)

	err = enc.Encode(CMap)
	if err != nil {
		return err, nil
	}

	return nil, &CMap.Timestamp
}

func DecodeData() (*scrapper.Currency, error) {

	var Curr *scrapper.Currency

	if _, err := os.Stat(CurrencyFile); os.IsNotExist(err) {
		err, _ = UpdateCurrency()
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(CurrencyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)

	err = dec.Decode(&Curr)
	if err != nil {
		return nil, err
	}

	return Curr, nil

}
