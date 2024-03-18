package repository

import (
	"fmt"

	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateRole(role models.Role) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description ) values ($1, $2) RETURNING id", rolesTable)
	row := r.db.QueryRow(query, role.Name, role.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, role_name, role_id ) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password, user.Role, user.RoleId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
