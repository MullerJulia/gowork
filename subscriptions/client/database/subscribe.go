package database

import (
	"context"
	"main/models"
)

// Subscribe creates inserts user subscription information into the database.
func (c *Client) Subscribe(ctx context.Context, s models.UserSubscription) error {
	stmt := "INSERT INTO subscriptions(user_id, subscription_id, status, charge_amount) VALUES ($1,$2,$3,$4) RETURNING subscription_id"

	_, err := c.DB.ExecContext(ctx, stmt, s.UserID, s.SubscriptionID, s.Status, s.ChargeAmount)
	if err != nil {
		return err
	}
	return nil
}
