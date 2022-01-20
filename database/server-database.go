package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"vendas/config"

	_ "github.com/denisenkom/go-mssqldb"
)

func (s *SQLStr) Connect() error {
	var err error
	if s.db, err = sql.Open("sqlserver", s.url.String()); err != nil {
		return err
	}
	return s.db.PingContext(context.Background())
}

//NewSQL ...
func NewSQL(conf *config.SQL) (*SQLStr, error) {
	s := &SQLStr{}
	return s, s.UpdateConfig(conf)
}

//UpdateConfig ...
func (s *SQLStr) UpdateConfig(conf *config.SQL) error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			return err
		}
	}
	s.conf = conf
	s.url = &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(conf.Username, conf.Password),
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		RawQuery: url.Values{"database": {conf.Db}}.Encode(),
	}
	return s.Connect()
}
