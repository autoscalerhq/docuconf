package docuconf

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

type testStruct struct {
	ConnectionString string `env:"CONNECTION_STRING"`
}

func TestLoadDotEnv(t *testing.T) {
	env, err := LoadDotEnv(filepath.Join("internal", "testing", "test.env"), testStruct{})
	require.Nil(t, err)
	require.Equal(t, "postgresql://user:password@localhost:5432/dbname", env.ConnectionString)
}
