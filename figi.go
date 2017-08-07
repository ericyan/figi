// Package figi implements the Financial Instrument Global Identifier.
package figi

import (
	"errors"
)

var (
	ErrInvalid = errors.New("invalid argument")
)

type FIGI [12]byte

func New(s string) (FIGI, error) {
	if len(s) != 12 {
		return FIGI{}, ErrInvalid
	}

	var id [12]byte
	var checksum int
	for i := 0; i < 12; i++ {
		switch {
		case '0' <= s[i] && s[i] <= '9':
			id[i] = s[i] - '0'
		case 'a' <= s[i] && s[i] <= 'z':
			id[i] = s[i] - 'a' + 10
		case 'A' <= s[i] && s[i] <= 'Z':
			id[i] = s[i] - 'A' + 10
		default:
			return FIGI{}, ErrInvalid
		}

		// Calculate checksum using the Luhn algorithm
		num := int(id[i])
		if (i%2 != 0) && (i != 11) {
			num *= 2
		}
		checksum += (num / 10) + (num % 10)
	}

	if checksum%10 != 0 {
		return FIGI{}, ErrInvalid
	}

	return FIGI(id), nil
}

func (id FIGI) String() string {
	s := make([]byte, 12)
	for i, c := range id {
		if c < 10 {
			s[i] = c + '0'
		} else {
			s[i] = c - 10 + 'A'
		}
	}

	return string(s)
}
