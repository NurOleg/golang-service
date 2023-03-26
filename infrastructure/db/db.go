package db

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "myGoApp/infrastructure/config"
    "fmt"
)

func GetConn(ec *config.EnvConfig) (*sql.DB, error) {
    var err error
    var db *sql.DB

    dbAddr := fmt.Sprintf("%s:%s", ec.DBHost, ec.DBPort)

    cfg := mysql.Config{
        User:   ec.DBLogin,
        Passwd: ec.DBPass,
        Net:    "tcp",
        Addr:   dbAddr,
        DBName: ec.DBName,
        AllowNativePasswords: true,
        }

        db, err = sql.Open("mysql", cfg.FormatDSN())

        if err != nil {
            return nil, err
        }

        return db, err
}