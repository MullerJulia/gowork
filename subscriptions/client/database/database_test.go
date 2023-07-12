package database

import (
	"context"
	"fmt"
	"main/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnString(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		config, err := config.LoadConfig()
		require.NoError(t, err)
		connString := ConnString(&config)
		varifString := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"localhost", 5432, "postgres", "postgres", "subscriptions")
		assert.Equal(t, connString, varifString)
	})
}

func TestInit(t *testing.T) {
	t.Run("initDB", func(t *testing.T) {
		var dbClient Client
		config, err := config.LoadConfig()
		require.NoError(t, err)
		ctx := context.Background()
		err = dbClient.Init(ctx, &config)
		require.NoError(t, err)
	})
}
