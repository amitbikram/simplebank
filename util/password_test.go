package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	pass := RandomString(6)
	s, err := HashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, s)
	err = CheckPassword(s, pass)
	require.NoError(t, err)
	err = CheckPassword(s, "pass")
	require.Error(t, err)
}
