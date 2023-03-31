package repository

import (
    "database/sql"
    "errors"
    "myGoApp/models"
)

type UserRepo struct {
    conn sql.DB
}

func NewRepo(conn *sql.DB) *UserRepo {
    return &UserRepo{
        conn: *conn,
    }
}

func (up *UserRepo) GetUserById(id int32) *models.User {
    var user models.User

    row := up.conn.QueryRow("SELECT email, password FROM users WHERE id = ? LIMIT 1", id)

    if err := row.Scan(&user.Email); err != nil {
        if errors.Is(sql.ErrNoRows, err) {
            return nil
        }
        panic(err)
    }

    return &user
}

func (up *UserRepo) GetUserByEmail(email string) *models.User {
    var user models.User

    row := up.conn.QueryRow("SELECT email FROM users WHERE email = ? LIMIT 1", email)

    if err := row.Scan(&user.Email); err != nil {
        if errors.Is(sql.ErrNoRows, err) {
            return nil
        }
        panic(err)
    }

    return &user
}