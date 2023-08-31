package repo

import (
	"context"
	"database/sql"
)

type DBHandler struct {
	DB  *sql.DB
	Ctx context.Context
}
