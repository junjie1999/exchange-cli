package calc

import (
	"errors"
	"time"

	"github.com/junjie1999/exchange-cli/gobutil"
	"github.com/junjie1999/exchange-cli/unit"
)

type Exchange struct {
	Base, Target string
	Amount       float64
	ExchangeRate float64
	Timestamp    int64
	Result       float64
}

var ExchangeRates gobutil.ExchangeRates

func NewExchange(base, target string, amount float64) (*Exchange, error) {

	Currencies, err := gobutil.DecodeData()
	if err != nil {
		return nil, err
	}
	return &Exchange{
		Base:         base,
		Target:       target,
		Amount:       amount,
		ExchangeRate: Currencies.ExchangeRates[base][target],
	}, nil
}

func (Ex *Exchange) convert() error {

	if !unit.XRateUnitExist(Ex.Base) {
		return errors.New("'" + Ex.Base + "' base currency not found")
	}

	if !unit.XRateUnitExist(Ex.Target) {
		return errors.New("'" + Ex.Target + "' target currency not found")
	}

	if Ex.ExchangeRate == -1 {
		Ex.Result = -1
		return errors.New("currency not found")
	}

	Ex.Result = Ex.Amount * Ex.ExchangeRate
	return nil
}

func Convert(amount float64, base, target string) (*Exchange, error) {
	Ex, err := NewExchange(base, target, amount)
	if err != nil {
		return nil, err
	}

	err = Ex.convert()
	if err != nil {
		return nil, err
	}

	return Ex, nil
}

func GetTimeStamp() string {
	t := time.Unix(ExchangeRates.Timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func init() {
	if ExchangeRates == nil {
		ExchangeRates, _ = gobutil.DecodeData()
	}
}
