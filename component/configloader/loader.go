// Package configloader decodes plain or age-encrypted configuration data.
package configloader

import (
	"fmt"

	"github.com/metacubex/mihomo/common/yaml"
	"github.com/metacubex/mihomo/component/age"
)

func Unmarshal(data []byte, target any, secretKeys ...string) error {
	plainData, err := Decrypt(data, secretKeys...)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(plainData, target)
}

func Decrypt(data []byte, secretKeys ...string) ([]byte, error) {
	plainData, err := age.DecryptBytes(data, secretKeys...)
	if err != nil {
		return nil, fmt.Errorf("decrypt config error: %w", err)
	}

	return plainData, nil
}
