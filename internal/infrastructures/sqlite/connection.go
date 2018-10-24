package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yojkim/math-test-app/internal/domains"
)

var db *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "tmp/db/math.db")
	if err != nil {
		panic(err.Error())
	}

	db.Table("problems").CreateTable(&domains.Problem{})
	db.Table("results").CreateTable(&domains.Result{})

	return db
}

func Close() {
	defer db.Close()
}
