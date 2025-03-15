package codecs

import (
	"testing"
)

var testCasesDecoder = []struct {
	encoding      string
	input         string
	expectedBytes []byte
}{
	{"utf-8", "Hello 世界", []byte("Hello 世界")},
	{"windows-1250", "Česká", []byte{0xC8, 0x65, 0x73, 0x6B, 0xE1}},
	{"windows-1251", "Привет", []byte{0xCF, 0xF0, 0xE8, 0xE2, 0xE5, 0xF2}},
	{"windows-1252", "Café", []byte{0x43, 0x61, 0x66, 0xE9}},
	{"windows-1253", "Γειά", []byte{0xC3, 0xE5, 0xE9, 0xDC}},
	{"windows-1254", "Merhaba", []byte{0x4D, 0x65, 0x72, 0x68, 0x61, 0x62, 0x61}},
	{"iso-8859-1", "Olé", []byte{0x4F, 0x6C, 0xE9}},
	{"iso-8859-2", "Český", []byte{0xC8, 0x65, 0x73, 0x6B, 0xFD}},
	{"iso-8859-5", "Привет", []byte{0xCF, 0xF0, 0xE8, 0xE2, 0xE5, 0xF2}},
	{"iso-8859-7", "Γειά", []byte{0xC3, 0xE5, 0xE9, 0xDC}},
	{"us-ascii", "Hello", []byte("Hello")},
}

func TestDecodeEncodeText(t *testing.T) {
	for _, tc := range testCasesDecoder {
		decoded, err := TextDecode(tc.expectedBytes, tc.encoding, TextDecoderOptions{Fatal: true, IgnoreBOM: true}, TextDecodeOptions{Stream: true})
		if err != nil {
			t.Errorf("Decoding failed for %s: %v", tc.encoding, err)
		}
		encoded, err := TextEncode(decoded, tc.encoding)
		if err != nil {
			t.Errorf("Encoding failed for %s: %v", tc.encoding, err)
		}
		if string(encoded) != string(tc.expectedBytes) {
			t.Errorf("Mismatch for %s: expected %v, got %v", tc.encoding, tc.expectedBytes, encoded)
		}
	}
}

func TestDecoderOptions(t *testing.T) {
	for _, tc := range testCasesDecoder {
		encoded, err := TextEncode(tc.input, tc.encoding)
		if err != nil {
			t.Errorf("Encoding failed for %s: %v", tc.encoding, err)
		}
		options := TextDecoderOptions{Fatal: true, IgnoreBOM: true}
		decodeOptions := TextDecodeOptions{Stream: true}
		decoded, err := TextDecode(encoded, tc.encoding, options, decodeOptions)
		if err != nil {
			t.Errorf("Decoding failed for %s with options: %v", tc.encoding, err)
		}
		if decoded != tc.input {
			t.Errorf("Mismatch for %s with options: expected %q, got %q", tc.encoding, tc.input, decoded)
		}
	}
}
