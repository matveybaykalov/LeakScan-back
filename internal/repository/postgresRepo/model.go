package postgresrepo

import "github.com/uptrace/bun"

type Challenges struct {
	bun.BaseModel `bun:"table:challenge"`

	Challenge0 int `bun:"chg_0"`
	Challenge1 int `bun:"chg_1"`
	Challenge2 int `bun:"chg_2"`
	Challenge3 int `bun:"chg_3"`
	Challenge4 int `bun:"chg_4"`
	Challenge5 int `bun:"chg_5"`
	Challenge6 int `bun:"chg_6"`
	Challenge7 int `bun:"chg_7"`
}

type Hash struct {
	bun.BaseModel `bun:"table:hash"`

	Password   string `bun:"password"`
	Challenge0 string `bun:"chg_0"`
	Challenge1 string `bun:"chg_1"`
	Challenge2 string `bun:"chg_2"`
	Challenge3 string `bun:"chg_3"`
	Challenge4 string `bun:"chg_4"`
	Challenge5 string `bun:"chg_5"`
	Challenge6 string `bun:"chg_6"`
	Challenge7 string `bun:"chg_7"`
}
