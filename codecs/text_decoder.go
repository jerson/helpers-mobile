package codecs

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var supportedEncodings = map[string]encoding.Encoding{
	"utf-8":        unicode.UTF8,
	"windows-1250": charmap.Windows1250,
	"windows-1251": charmap.Windows1251,
	"windows-1252": charmap.Windows1252,
	"windows-1253": charmap.Windows1253,
	"windows-1254": charmap.Windows1254,
	"windows-1255": charmap.Windows1255,
	"windows-1256": charmap.Windows1256,
	"windows-1257": charmap.Windows1257,
	"windows-1258": charmap.Windows1258,
	"iso-8859-1":   charmap.ISO8859_1,
	"iso-8859-2":   charmap.ISO8859_2,
	"iso-8859-3":   charmap.ISO8859_3,
	"iso-8859-4":   charmap.ISO8859_4,
	"iso-8859-5":   charmap.ISO8859_5,
	"iso-8859-6":   charmap.ISO8859_6,
	"iso-8859-7":   charmap.ISO8859_7,
	"iso-8859-8":   charmap.ISO8859_8,
	"iso-8859-9":   charmap.ISO8859_9,
	"iso-8859-10":  charmap.ISO8859_10,
	"iso-8859-14":  charmap.ISO8859_14,
	"iso-8859-15":  charmap.ISO8859_15,
	"iso-8859-16":  charmap.ISO8859_16,
	"us-ascii":     charmap.ISO8859_1,
}

type TextDecoderOptions struct {
	Fatal     bool
	IgnoreBOM bool
}

type TextDecodeOptions struct {
	Stream bool
}

func TextDecode(input []byte, encoding string, options TextDecoderOptions, decodeOptions TextDecodeOptions) (string, error) {
	enc, ok := supportedEncodings[encoding]
	if !ok {
		return "", errors.New("unsupported encoding: " + encoding)
	}

	if options.IgnoreBOM {
		input = bytes.TrimPrefix(input, []byte{0xEF, 0xBB, 0xBF})
	}

	reader := transform.NewReader(bytes.NewReader(input), enc.NewDecoder())

	if decodeOptions.Stream {
		buffer := make([]byte, 1024)
		var output bytes.Buffer
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				output.Write(buffer[:n])
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				if options.Fatal {
					return "", err
				}
				break
			}
		}
		return output.String(), nil
	}

	decoded, err := io.ReadAll(reader)
	if err != nil && options.Fatal {
		return "", err
	}

	return string(decoded), nil
}
