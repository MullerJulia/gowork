// Package database contains a Postgres client and methods for communicating with the database.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"main/config"
)

// Client holds the database client and prepared statements.
type Client struct {
	DB *sql.DB
}

func ConnString(config *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DatabaseUrl, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseDb)

}

// Init sets up a new database client.
func (c *Client) Init(ctx context.Context, config *config.Config) error {
	db, err := sql.Open("postgres", ConnString(config))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	c.DB = db
	return nil
}

// Close closes the database connection and statements.
func (c *Client) Close() error {
	if err := c.DB.Close(); err != nil {
		return err
	}
	return nil
}
