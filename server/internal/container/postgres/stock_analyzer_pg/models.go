// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package stock_analyzer_pg

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Author struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}