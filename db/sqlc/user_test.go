package db

import (
	"context"
	"testing"
	"time"

	"github.com/sarangi/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) Users {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	userParam := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), userParam)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, userParam.Username, user.Username)
	require.Equal(t, userParam.FullName, user.FullName)
	require.Equal(t, userParam.Email, user.Email)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)
	assert.NoError(t, err)
	assert.NotEmpty(t, user2)
	assert.Equal(t, user.Username, user2.Username)
	assert.Equal(t, user.FullName, user2.FullName)
	assert.Equal(t, user.Email, user2.Email)
	assert.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)
}
