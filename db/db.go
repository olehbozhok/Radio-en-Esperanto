package botdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitDB(dbstring string, debug bool) (*sql.DB, error) {

	db, err := sql.Open("mysql", dbstring)
	if err != nil {
		return nil, err
	}
	boil.SetDB(db)
	boil.DebugMode = debug
	err = db.Ping()

	return db, err
}
