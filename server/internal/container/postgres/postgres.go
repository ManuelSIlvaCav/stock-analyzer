package postgres

import (
	"context"
	"fmt"
	"os"
	"stockanalyzer/internal/container/config"

	"github.com/jackc/pgx/v5"
)

type PostgresSQL interface {
	GetConnection() *pgx.Conn
}

type postgresSQL struct {
	connection *pgx.Conn
}

func NewPostgresSQL(config config.PostgresConfig) PostgresSQL {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", config.Host, config.User, config.Password, config.DBName, config.Port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//defer conn.Close(ctx)

	return &postgresSQL{
		connection: conn,
	}
}

func (p *postgresSQL) GetConnection() *pgx.Conn {
	return p.connection
}
