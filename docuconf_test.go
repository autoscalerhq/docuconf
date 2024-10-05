package docuconf

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

type testStruct struct {
	ConnectionString string `env:"CONNECTION_STRING"`
	SomeBool         bool   `env:"SOME_BOOL"`
	SomeOtherBool    bool   `env:"SOME_OTHER_BOOL"`
}

func TestLoadDotEnv(t *testing.T) {
	env, err := LoadDotEnv(filepath.Join("internal", "testing", "test.env"), testStruct{})
	require.Nil(t, err)
	require.Equal(t, "postgresql://user:password@localhost:5432/dbname", env.ConnectionString)
	require.Equal(t, true, env.SomeBool)
	require.Equal(t, false, env.SomeOtherBool)
}
