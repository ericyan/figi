package figi

import (
	"fmt"
	"testing"
)

func TestTextMarshalling(t *testing.T) {
	cases := []struct {
		in  []byte
		out []byte
		err error
	}{
		{nil, nil, ErrInvalid},
		{[]byte("ABC012345678"), nil, ErrInvalid},
		{[]byte("bbg000blnq16"), []byte("BBG000BLNQ16"), nil},
	}

	for _, c := range cases {
		id := new(FIGI)
		err := id.UnmarshalText(c.in)
		if (c.out != nil && err != nil) || (c.out == nil && err != c.err) {
			t.Errorf("unexpected error: %s", err)
		}

		out, err := id.MarshalText()
		if err != nil && err != c.err {
			t.Errorf("unexpected error: %s", err)
		}
		for i := 0; i < len(c.out); i++ {
			if out[i] != c.out[i] {
				t.Errorf("unexpected result: got %s, want %s", out, c.out)
			}
		}
	}
}

func ExampleClient() {
	req := []MappingRequest{
		{
			IDType:  "ID_EXCH_SYMBOL",
			IDValue: "5",
			MIC:     "XHKG",
		},
		{
			IDType:  "ID_SEDOL",
			IDValue: "3070732",
		},
		{
			IDType:  "ID_ISIN",
			IDValue: "XS0202077953",
		},
		{
			IDType:  "ID_CUSIP",
			IDValue: "BOGUSCUSIP",
		},
	}

	c := NewClient()
	resp, err := c.Query(req)
	if err != nil {
		fmt.Println(err)
	}

	for i, mapping := range resp {
		if mapping.Success() {
			for _, result := range mapping.Data {
				fmt.Printf("%s => %s: %s\n", req[i].IDValue, result.FIGI, result.Name)
			}
		} else {
			fmt.Printf("%s => ERROR: %s\n", req[i].IDValue, mapping.Error)
		}
	}

	// Output:
	// 5 => BBG000BQ9HJ2: HSBC HOLDINGS PLC
	// 3070732 => BBG000BT7PY3: BANK OF IRELAND
	// 3070732 => BBG000FXQJP3: BANK OF IRELAND
	// 3070732 => BBG000G7GBS4: BANK OF IRELAND
	// 3070732 => BBG000JP1BD3: BANK OF IRELAND
	// 3070732 => BBG0025BSC99: BANK OF IRELAND
	// XS0202077953 => BBG0000585F1: WAL-MART STORES INC
	// BOGUSCUSIP => ERROR: No identifier found.
}
