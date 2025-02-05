package postgres

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config interface {
	GetHost() string
	GetPort() string
	GetDatabase() string
	GetUser() string
	GetPassword() string
	GetSSLMode() string
}

type ConfigModel struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMode  string
}

func (c ConfigModel) GetHost() string {
	return c.Host
}

func (c ConfigModel) GetPort() string {
	return c.Port
}

func (c ConfigModel) GetPassword() string {
	return c.Password
}

func (c ConfigModel) GetDatabase() string {
	return c.Database
}

func (c ConfigModel) GetUser() string {
	return c.User
}

func (c ConfigModel) GetSSLMode() string {
	return c.SSLMode
}

func InitSqlDB(cfg Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetHost(),
		cfg.GetPort(),
		cfg.GetUser(),
		cfg.GetPassword(),
		cfg.GetDatabase(),
		cfg.GetSSLMode(),
	)

	database, err := sqlx.Open("pgx", connectionURL)
	if err != nil {
		return nil, err
	}

	// database.SetMaxOpenConns(cfg.Postgres.Settings.MaxOpenConns)
	// database.SetConnMaxIdleTime(cfg.Postgres.Settings.MaxIdletime * time.Second)
	// database.SetConnMaxLifetime(cfg.Postgres.Settings.MaxLifetime * time.Second)
	// database.SetMaxIdleConns(cfg.Postgres.Settings.MaxIdleConns)

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
