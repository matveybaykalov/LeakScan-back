package clickhouserepo

import "github.com/uptrace/go-clickhouse/ch"

type ClickHouseRepo struct {
	con *ch.DB
}

func New(con *ch.DB) *ClickHouseRepo {
	return &ClickHouseRepo{
		con: con,
	}
}
