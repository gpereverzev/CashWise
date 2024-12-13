package handlers

import (
	"github.com/gpereverzev/CashWise/src/data"
	"net/http"
)

const MasterDBContextKey = iota

func PostgresDB(r *http.Request) data.MasterDB {
	return r.Context().Value(MasterDBContextKey).(data.MasterDB)
}
