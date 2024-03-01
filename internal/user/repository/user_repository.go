package repository

import (
	"database/sql"
	"fmt"

	"github.com/Patrikesm/basic-solid-api/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	_, err := repo.DB.Exec(
		"INSERT INTO users (id, name, status, age) VALUES ?, ?, ?, ?",
		user.ID, user.Name, user.Status, user.Age,
	)

	if err != nil {
		fmt.Println("Deu pau:", err.Error())
		return err
	}

	return nil
}

func (repo *UserRepository) ListAll() ([]*models.User, error) {
	rows, err := repo.DB.Query("SELECT * FROM users")

	fmt.Println(rows)

	if err != nil {
		return nil, err
	}

	var userList []*models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Status,
			&user.Age,
		)

		if err != nil {
			return nil, err
		}

		userList = append(userList, &user)
	}

	return userList, nil
}
