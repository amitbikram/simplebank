package util

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case USD, INR, EUR:
		return true
	}
	return false
}
