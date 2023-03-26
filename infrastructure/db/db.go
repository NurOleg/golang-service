package db

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

func GetConn() (*sql.DB, error) {
    var err error
    var db *sql.DB

    cfg := mysql.Config{
        User:   "root",
        Passwd: "1234567",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "modelshub",
        AllowNativePasswords: true,
        }

        db, err = sql.Open("mysql", cfg.FormatDSN())

        if err != nil {
            return nil, err
        }

        return db, err
}