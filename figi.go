// Package figi implements the Financial Instrument Global Identifier.
package figi

import (
	"errors"
)

var (
	ErrInvalid = errors.New("invalid argument")
)

type FIGI [12]byte

func (id *FIGI) UnmarshalText(text []byte) error {
	if len(text) != 12 {
		return ErrInvalid
	}

	var checksum int
	for i := 0; i < 12; i++ {
		c := text[i]
		switch {
		case '0' <= c && c <= '9':
			id[i] = c - '0'
		case 'a' <= c && c <= 'z':
			id[i] = c - 'a' + 10
		case 'A' <= c && c <= 'Z':
			id[i] = c - 'A' + 10
		default:
			return ErrInvalid
		}

		// Calculate checksum using the Luhn algorithm
		num := int(id[i])
		if (i%2 != 0) && (i != 11) {
			num *= 2
		}
		checksum += (num / 10) + (num % 10)
	}

	if checksum%10 != 0 {
		return ErrInvalid
	}

	return nil
}

func (id *FIGI) MarshalText() ([]byte, error) {
	s := make([]byte, 12)
	for i, c := range id {
		if c < 10 {
			s[i] = c + '0'
		} else {
			s[i] = c - 10 + 'A'
		}
	}

	return s, nil
}

func (id *FIGI) String() string {
	text, err := id.MarshalText()
	if err != nil {
		return "<invalid>"
	}

	return string(text)
}
