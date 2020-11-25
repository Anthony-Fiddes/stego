package stego

import "testing"

func TestSetSuffix(t *testing.T) {
	tests := []struct {
		name         string
		x            byte
		suffix       byte
		suffixLength byte
		expected     byte
	}{
		{
			name:         "FF to FA",
			x:            0xFF,
			suffix:       0x0A,
			suffixLength: 4,
			expected:     0xFA,
		},
		{
			name:         "00000000 to 00000001",
			x:            0x0,
			suffix:       0b00000001,
			suffixLength: 1,
			expected:     0b00000001,
		},
		{
			name:         "00001111 to 00001101",
			x:            0x0F,
			suffix:       0b00000101,
			suffixLength: 3,
			expected:     0b00001101,
		},
		{
			name:         "00001010 to 00001011",
			x:            0b00001010,
			suffix:       0b00000011,
			suffixLength: 3,
			expected:     0b00001011,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := setSuffix(test.x, test.suffix, test.suffixLength)
			if result != test.expected {
				t.Fatalf(
					"Ran setSuffix(%b, %b, %d) expecting a result of %08b, instead received %08b\n",
					test.x,
					test.suffix,
					test.suffixLength,
					test.expected,
					result,
				)
			}
		})
	}
}
