package db_service

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var ErrNotFound = fmt.Errorf("record not found")
var ErrConflict = fmt.Errorf("record already exists")

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

type PostgresService struct {
	DB *sql.DB
}

func NewPostgresService(cfg PostgresConfig) (*PostgresService, error) {

	env := func(name string, defaultValue string) string {
		if value, ok := os.LookupEnv(name); ok {
			return value
		}
		return defaultValue
	}

	if cfg.Host == "" {
		cfg.Host = env("POSTGRES_HOST", "localhost")
	}

	if cfg.Port == 0 {
		port, err := strconv.Atoi(env("POSTGRES_PORT", "5432"))
		if err == nil {
			cfg.Port = port
		} else {
			cfg.Port = 5432
		}
	}

	if cfg.User == "" {
		cfg.User = env("POSTGRES_USER", "postgres")
	}

	if cfg.Password == "" {
		cfg.Password = env("POSTGRES_PASSWORD", "postgres")
	}

	if cfg.Database == "" {
		cfg.Database = env("POSTGRES_DB", "ambulance")
	}

	if cfg.SSLMode == "" {
		cfg.SSLMode = env("POSTGRES_SSLMODE", "disable")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn) //connect to database
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &PostgresService{
		DB: db,
	}, nil
}

func (s *PostgresService) Close() error {
	if s.DB == nil {
		return nil
	}
	return s.DB.Close()
}

func (s *PostgresService) Ping(ctx context.Context) error {
	return s.DB.PingContext(ctx)
}