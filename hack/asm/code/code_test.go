package code

import "testing"

func TestEncodeComp(t *testing.T) {
	for _, test := range []struct {
		mnemonic string
		expected string
	}{
		{},
	} {
		if actual := Comp(test.mnemonic); actual != test.expected {
			t.Errorf("%s should be encoded to %s, got %s", test.mnemonic, test.expected, actual)
		}
	}
}

func TestEncodeDest(t *testing.T) {
	for _, test := range []struct {
		mnemonic string
		expected string
	}{
		{},
	} {
		if actual := Dest(test.mnemonic); actual != test.expected {
			t.Errorf("%s should be encoded to %s, got %s", test.mnemonic, test.expected, actual)
		}
	}
}

func TestEncodeJump(t *testing.T) {
	for _, test := range []struct {
		mnemonic string
		expected string
	}{
		{},
	} {
		if actual := Jump(test.mnemonic); actual != test.expected {
			t.Errorf("%s should be encoded to %s, got %s", test.mnemonic, test.expected, actual)
		}
	}
}
