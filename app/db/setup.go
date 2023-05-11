package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"sync"
)

var (
	mysql  *sqlx.DB
	m      sync.Mutex
	driver = "mysql"
)

func MySQL() (error, *sqlx.DB) {
	m.Lock()
	defer m.Unlock()

	if mysql != nil {
		return nil, mysql
	}

	fmt.Println("init.")
	db, err := sqlx.Open(driver, viper.GetString("rds.mysql.dsn"))
	if err != nil {
		return err, nil
	}

	if err := db.Ping(); err != nil {
		return err, nil
	}

	db.SetMaxIdleConns(viper.GetInt("rds.mysql.idle"))
	db.SetMaxOpenConns(viper.GetInt("rds.mysql.open"))
	mysql = db

	return nil, mysql
}

func Close() error {
	return CloseMySQL()
}

func CloseMySQL() error {
	if mysql != nil {
		return mysql.Close()
	}

	return nil
}
