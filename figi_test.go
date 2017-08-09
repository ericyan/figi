package figi

import "testing"

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
