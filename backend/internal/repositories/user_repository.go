package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"myproject/backend/internal/models"

	"github.com/Masterminds/squirrel"
)




type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAll() ([]models.User, error) {
    query := squirrel.Select("id", "user_name", "first_name", "last_name", "email", "user_status", "department").
        From("users").
        PlaceholderFormat(squirrel.Dollar)

    sql, args, err := query.ToSql()
    if err != nil {
        return nil, err
    }

    rows, err := r.db.Query(sql, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(
            &user.ID,
            &user.UserName,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.UserStatus,
            &user.Department,
        ); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}


func (r *UserRepositoryImpl) GetByID(id int) (*models.User, error) {
    query := squirrel.Select("id", "user_name", "first_name", "last_name", "email", "user_status", "department").
        From("users").
        Where(squirrel.Eq{"id": id}).
        PlaceholderFormat(squirrel.Dollar)

    sql, args, err := query.ToSql()
    if err != nil {
        log.Printf("Error building SQL: %v", err)
        return nil, err
    }

    log.Printf("SQL Query: %s", sql)
    log.Printf("SQL Args: %v", args)

    var user models.User
    err = r.db.QueryRow(sql, args...).Scan(
        &user.ID,
        &user.UserName,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.UserStatus,
        &user.Department,
    )
    if err != nil {
        if err.Error() == "sql: no rows in result set" {
            log.Printf("No user found with ID %d", id)
            return nil, nil
        }
        log.Printf("Error executing SQL: %v", err)
        return nil, err
    }

    return &user, nil
}




func (r *UserRepositoryImpl) Create(user *models.User) error {
    query := squirrel.Insert("users").
        Columns("user_name", "first_name", "last_name", "email", "user_status", "department").
        Values(user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department).
        Suffix("RETURNING id").
        PlaceholderFormat(squirrel.Dollar)

    sql, args, err := query.ToSql()
    if err != nil {
        log.Printf("Error building SQL: %v", err)
        return err
    }

    log.Printf("SQL Query: %s", sql)
    log.Printf("SQL Args: %v", args)

    err = r.db.QueryRow(sql, args...).Scan(&user.ID)
    if err != nil {
        log.Printf("Error executing SQL: %v", err)
        return err
    }

    return nil
}



func (r *UserRepositoryImpl) Update(user *models.User) error {
    query := squirrel.Update("users").
        Set("user_name", user.UserName).
        Set("first_name", user.FirstName).
        Set("last_name", user.LastName).
        Set("email", user.Email).
        Set("user_status", user.UserStatus).
        Set("department", user.Department).
        Where(squirrel.Eq{"id": user.ID}).
        PlaceholderFormat(squirrel.Dollar)

    sql, args, err := query.ToSql()
    if err != nil {
        log.Printf("Error building SQL: %v", err)
        return err
    }

    log.Printf("Update SQL Query: %s", sql)
    log.Printf("Update SQL Args: %v", args)

    result, err := r.db.Exec(sql, args...)
    if err != nil {
        log.Printf("Error executing update SQL: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("Error getting rows affected: %v", err)
        return err
    }

    log.Printf("Rows affected by update: %d", rowsAffected)

    if rowsAffected == 0 {
        return fmt.Errorf("no user found with ID %d", user.ID)
    }

    return nil
}


func (r *UserRepositoryImpl) Delete(id int) error {
    query := squirrel.Delete("users").
        Where(squirrel.Eq{"id": id}).
        PlaceholderFormat(squirrel.Dollar)

    sql, args, err := query.ToSql()
    if err != nil {
        log.Printf("Error building SQL: %v", err)
        return err
    }

    log.Printf("Delete SQL Query: %s", sql)
    log.Printf("Delete SQL Args: %v", args)

    result, err := r.db.Exec(sql, args...)
    if err != nil {
        log.Printf("Error executing delete SQL: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("Error getting rows affected: %v", err)
        return err
    }

    log.Printf("Rows affected by delete: %d", rowsAffected)

    if rowsAffected == 0 {
        return fmt.Errorf("no user found with ID %d", id)
    }

    return nil
}


