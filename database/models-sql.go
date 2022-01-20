package database

import (
	"database/sql"
	"net/url"
	"vendas/config"
)

// SQLStr ...
type SQLStr struct {
	conf *config.SQL
	url  *url.URL
	db   *sql.DB
}
