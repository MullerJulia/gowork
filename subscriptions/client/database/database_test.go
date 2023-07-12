package database

import (
	"fmt"
	"main/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnString(t *testing.T) {
	t.Run("congig", func(t *testing.T) {
		config, err := config.LoadConfig()
		require.NoError(t, err)
		connString := ConnString(&config)
		varifString := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"localhost", 5432, "postgres", "postgres", "subscriptions")
		assert.Equal(t, connString, varifString)
	})
}
