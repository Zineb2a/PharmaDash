package util

import (
	"errors"
	"regexp"
)

// ValidateCardInfo checks if the provided card information has a valid format
func ValidateCardInfo(cardNumber string, expiryDate string, cvv string) error {
	cardNumberPattern := `^\d{16}$`
	expiryDatePattern := `^(0[1-9]|1[0-2])\/\d{2}$` // MM/YY
	cvvPattern := `^\d{3}$`

	if !regexp.MustCompile(cardNumberPattern).MatchString(cardNumber) {
		return errors.New("invalid card number format")
	}
	if !regexp.MustCompile(expiryDatePattern).MatchString(expiryDate) {
		return errors.New("invalid expiry date format")
	}
	if !regexp.MustCompile(cvvPattern).MatchString(cvv) {
		return errors.New("invalid CVV format")
	}
	return nil
}
