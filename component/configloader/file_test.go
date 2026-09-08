package configloader

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.yaml")
	data := []byte("mixed-port: 7890\n")
	require.NoError(t, os.WriteFile(path, data, 0o600))

	actual, err := ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, data, actual)
}

func TestReadFileMissing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.yaml")

	_, err := ReadFile(path)
	require.Error(t, err)
	assert.True(t, errors.Is(err, fs.ErrNotExist))
}

func TestReadFileEmpty(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.yaml")
	require.NoError(t, os.WriteFile(path, nil, 0o600))

	_, err := ReadFile(path)
	require.EqualError(t, err, "configuration file "+path+" is empty")
}

func TestReadFileDirectory(t *testing.T) {
	path := t.TempDir()

	_, err := ReadFile(path)
	require.Error(t, err)
}
