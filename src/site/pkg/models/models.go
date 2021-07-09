package models

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

type Webhook struct {
	Id int64 `db:"id"`
	AccountId  int64 `db:"account_id"`
	Domain  string `db:"domain"`
	ProductSysname  string `db:"product_sysname"`
	Request types.JSONText `json:"request" db:"request"`
	Status  string `db:"status"`
	Error  types.JSONText `json:"error" db:"error"`
	Info  types.JSONText `json:"info" db:"info"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}