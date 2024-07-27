package unit

var X_Rate_CountryUnits = map[string]string{
	"ARS": "Argentine Peso",
	"AUD": "Australian Dollar",
	"BHD": "Bahraini Dinar",
	"BWP": "Botswana Pula",
	"BRL": "Brazilian Real",
	"BND": "Bruneian Dollar",
	"BGN": "Bulgarian Lev",
	"CAD": "Canadian Dollar",
	"CLP": "Chilean Peso",
	"CNY": "Chinese Yuan Renminbi",
	"COP": "Colombian Peso",
	"CZK": "Czech Koruna",
	"DKK": "Danish Krone",
	"EUR": "Euro",
	"HKD": "Hong Kong Dollar",
	"HUF": "Hungarian Forint",
	"ISK": "Icelandic Krona",
	"INR": "Indian Rupee",
	"IDR": "Indonesian Rupiah",
	"IRR": "Iranian Rial",
	"ILS": "Israeli Shekel",
	"JPY": "Japanese Yen",
	"KZT": "Kazakhstani Tenge",
	"KRW": "South Korean Won",
	"KWD": "Kuwaiti Dinar",
	"LYD": "Libyan Dinar",
	"MYR": "Malaysian Ringgit",
	"MUR": "Mauritian Rupee",
	"MXN": "Mexican Peso",
	"NPR": "Nepalese Rupee",
	"NZD": "New Zealand Dollar",
	"NOK": "Norwegian Krone",
	"OMR": "Omani Rial",
	"PKR": "Pakistani Rupee",
	"PHP": "Philippine Peso",
	"PLN": "Polish Zloty",
	"QAR": "Qatari Riyal",
	"RON": "Romanian New Leu",
	"RUB": "Russian Ruble",
	"SAR": "Saudi Arabian Riyal",
	"SGD": "Singapore Dollar",
	"ZAR": "South African Rand",
	"LKR": "Sri Lankan Rupee",
	"SEK": "Swedish Krona",
	"CHF": "Swiss Franc",
	"TWD": "Taiwan New Dollar",
	"THB": "Thai Baht",
	"TTD": "Trinidadian Dollar",
	"TRY": "Turkish Lira",
	"AED": "Emirati Dirham",
	"GBP": "British Pound",
	"USD": "US Dollar",
	"VEF": "Venezuelan Bolivar",
}

var X_RateUnits []string

func init() {
	initXRateUnit()
}

func initXRateUnit() {
	for unit := range X_Rate_CountryUnits {
		X_RateUnits = append(X_RateUnits, unit)
	}
}

func GetXRateUnit(country string) string {
	for key, value := range X_Rate_CountryUnits {
		if value == country {
			return key
		}
	}
	return ""
}

func XRateUnitExist(unit string) bool {
	for _, u := range X_RateUnits {
		if u == unit {
			return true
		}
	}
	return false
}
