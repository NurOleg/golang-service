package db

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "myGoApp/infrastructure/helpers"
    "fmt"
)

func GetConn() (*sql.DB, error) {
    var err error
    var db *sql.DB

    dbAddr := fmt.Sprintf("%s:%s", helpers.GetEnvVar("DB_HOST"), helpers.GetEnvVar("DB_PORT"))

    cfg := mysql.Config{
        User:   helpers.GetEnvVar("DB_LOGIN"),
        Passwd: helpers.GetEnvVar("DB_PASS"),
        Net:    "tcp",
        Addr:   dbAddr,
        DBName: helpers.GetEnvVar("DB_NAME"),
        AllowNativePasswords: true,
        }

        db, err = sql.Open("mysql", cfg.FormatDSN())

        if err != nil {
            return nil, err
        }

        return db, err
}