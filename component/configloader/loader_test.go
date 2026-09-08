package configloader

import (
	"testing"

	"github.com/metacubex/mihomo/component/age"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfig struct {
	Port int `yaml:"port"`
	DNS  struct {
		Enable bool `yaml:"enable"`
	} `yaml:"dns"`
}

func TestUnmarshalPlainConfig(t *testing.T) {
	var cfg testConfig
	err := Unmarshal([]byte("port: 7890\ndns:\n  enable: true\n"), &cfg)
	require.NoError(t, err)

	assert.Equal(t, 7890, cfg.Port)
	assert.True(t, cfg.DNS.Enable)
}

func TestUnmarshalEncryptedConfig(t *testing.T) {
	secretKey, publicKey, err := age.GenX25519KeyPair()
	require.NoError(t, err)

	encrypted, err := age.EncryptBytes([]byte("port: 7890\n"), publicKey)
	require.NoError(t, err)

	var cfg testConfig
	err = Unmarshal(encrypted, &cfg, secretKey)
	require.NoError(t, err)
	assert.Equal(t, 7890, cfg.Port)
}

func TestUnmarshalDecryptError(t *testing.T) {
	_, publicKey, err := age.GenX25519KeyPair()
	require.NoError(t, err)

	encrypted, err := age.EncryptBytes([]byte("port: 7890\n"), publicKey)
	require.NoError(t, err)

	var cfg testConfig
	err = Unmarshal(encrypted, &cfg)
	require.Error(t, err)
	assert.ErrorContains(t, err, "decrypt config error")
}

func TestDecryptPlainConfigReturnsOriginalData(t *testing.T) {
	data := []byte("port: 7890\n")

	decrypted, err := Decrypt(data)
	require.NoError(t, err)
	assert.Equal(t, data, decrypted)
}
