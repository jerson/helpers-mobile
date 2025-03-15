package codecs

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/text/transform"
)

func TextEncode(input string, encoding string) ([]byte, error) {
	enc, ok := supportedEncodings[encoding]
	if !ok {
		return nil, errors.New("unsupported encoding: " + encoding)
	}

	reader := transform.NewReader(bytes.NewReader([]byte(input)), enc.NewEncoder())
	encoded, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}
