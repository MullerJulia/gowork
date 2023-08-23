package database

import (
	"context"
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
		ctx            context.Context
	}{
		{
			name: "success",
			args: models.UserSubscription{
				UserID:         "1",
				SubscriptionID: "1",
				ChargeAmount:   1,
				Status:         "active",
			},
			err: nil,
			ctx: context.Background(),
		},
	}

	db, dbClose := DbHelper(t)
	defer dbClose()

	SetUp(t, db, nil)
	t.Cleanup(func() {
	})

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			s := Client{DB: db}
			err := s.Subscribe(tc.ctx, tc.args)
			require.NoError(t, err)

			sqlStatement := "SELECT user_id, subscription_id, status, charge_amount FROM subscriptions WHERE user_id = $1"
			rows, err := db.Query(sqlStatement, tc.args.UserID)
			require.NoError(t, err)
			defer rows.Close()

			dbSubscriptions := models.UserSubscription{}
			for rows.Next() {
				dbSubscription := models.UserSubscription{}
				err := rows.Scan(&dbSubscription.UserID, &dbSubscription.SubscriptionID, &dbSubscription.Status, &dbSubscription.ChargeAmount)
				require.NoError(t, err)
			}

			assert.Equal(t, tc.args, dbSubscriptions)
		})
	}
}
