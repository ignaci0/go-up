package decorator

import (
	"github.com/ufoscout/go-up/reader"
	"strings"
)

type KeyStringReplacerDecoratorReader struct {
	Reader reader.Reader
	OldString string
	NewString string
}

func (f *KeyStringReplacerDecoratorReader) Read() (map[string]reader.Property, error) {

	result, err := f.Reader.Read()

	if err != nil {
		return nil, err
	}

	output := map[string]reader.Property{}

	for k, v := range result {
		output[strings.Replace(k, f.OldString, f.NewString, -1)] = v
	}

	return output, nil

}
