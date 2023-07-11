package database

import (
	"database/sql"
	"fmt"
	"main/models"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSubscriptions(t *testing.T) {
	db, dbClose := DbHelper(t)
	defer dbClose()

	SetUp(t, db, nil)
	t.Cleanup(func() {
	})

	t.Run("succes", func(t *testing.T) {
		s := Client{DB: db}

		subscriptions, err := s.GetSubscriptions()
		require.NoError(t, err)

		dbSubscriptions := []models.UserSubscription{}
		rows, err := db.Query("SELECT user_id, subscription_id, status, charge_amount FROM subscriptions")
		require.NoError(t, err)

		for rows.Next() {
			dbSubscription := models.UserSubscription{}
			err := rows.Scan(&dbSubscription.UserID, &dbSubscription.SubscriptionID, &dbSubscription.Status, &dbSubscription.ChargeAmount)
			require.NoError(t, err)
			dbSubscriptions = append(dbSubscriptions, dbSubscription)
		}
		assert.Equal(t, subscriptions, dbSubscriptions)
	})

}

func TestGetUserSubscriptions(t *testing.T) {
	test := []struct {
		name          string
		userID        string
		subscriptions []models.UserSubscription
		err           error
	}{
		{
			name:   "success",
			userID: "1",
			err:    nil,
		},
		{
			name:          "no user in database with id",
			userID:        "2",
			subscriptions: []models.UserSubscription([]models.UserSubscription(nil)),
			err:           fmt.Errorf("no user with id: %s", "2"),
		},
	}

	db, dbClose := DbHelper(t)
	defer dbClose()

	SetUp(t, db, nil)
	t.Cleanup(func() {
	})

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			sqlStatement := "SELECT user_id, subscription_id, status, charge_amount FROM subscriptions WHERE user_id = $1"
			rows, err := db.Query(sqlStatement, tc.userID)
			require.NoError(t, err)
			defer rows.Close()

			dbSubscriptions := []models.UserSubscription{}
			for rows.Next() {
				dbSubscription := models.UserSubscription{}
				err := rows.Scan(&dbSubscription.UserID, &dbSubscription.SubscriptionID, &dbSubscription.Status, &dbSubscription.ChargeAmount)
				require.NoError(t, err)
				dbSubscriptions = append(dbSubscriptions, dbSubscription)
			}

			s := Client{DB: db}
			subscriptions, err := s.GetUserSubscriptions(tc.userID)
			if tc.name != "success" {
				assert.Equal(t, tc.subscriptions, subscriptions)
				assert.Equal(t, tc.err, err)
			}
			if tc.name == "success" {
				assert.Equal(t, dbSubscriptions, subscriptions)
				assert.Equal(t, tc.err, err)
			}

		})

	}

}

func SetUp(t *testing.T, db *sql.DB, data any) {
	t.Helper()
}

func DbHelper(t *testing.T) (*sql.DB, func() error) {
	t.Helper()
	db, err := sql.Open("postgres", ConnString())
	require.NoError(t, err)

	return db, db.Close
}
