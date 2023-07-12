package database

import (
	"main/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscribe(t *testing.T) {
	test := []struct {
		name           string
		args           models.UserSubscription
		subscriptionID string
		err            error
	}{
		{
			name: "success",
			args: models.UserSubscription{
				UserID:         "1",
				SubscriptionID: "1",
				ChargeAmount:   1,
				Status:         "active",
			},
			subscriptionID: "1",
			err:            nil,
		},
	}

	db, dbClose := DbHelper(t)
	defer dbClose()

	SetUp(t, db, nil)
	t.Cleanup(func() {
	})

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			sqlStatement := "INSERT INTO subscriptions(user_id, subscription_id, status, charge_amount) VALUES ($1,$2,$3,$4) RETURNING subscription_id"
			err := db.QueryRow(sqlStatement, tc.args.UserID, tc.args.SubscriptionID, tc.args.Status, tc.args.ChargeAmount).Scan(&tc.subscriptionID)
			require.NoError(t, err)

			s := Client{DB: db}
			dbSubscriptionID, err := s.Subscribe(tc.args)
			assert.Equal(t, tc.subscriptionID, dbSubscriptionID)
			require.NoError(t, err)
		})
	}
}
