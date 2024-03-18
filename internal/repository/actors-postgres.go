package repository

import (
	"fmt"

	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/jmoiron/sqlx"
)

type ActorsPostgres struct {
	db *sqlx.DB
}

func NewActorsPostgres(db *sqlx.DB) *ActorsPostgres {
	return &ActorsPostgres{db: db}
}

func (r *ActorsPostgres) CreateActor(actor models.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, gender, date_birth, movie_id) values ($1, $2, $3, $4) RETURNING id", actorsTable)
	row := r.db.QueryRow(query, actor.Name, actor.Gender, actor.Date_birth, actor.Movie_id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActorsPostgres) GetActorById(id int) (models.Actor, error) {
	var actor models.Actor
	query := fmt.Sprintf("SELECT name, gender, date_birth, movie_id FROM %s WHERE id=$1", actorsTable)
	err := r.db.Get(&actor, query, id)
	return actor, err
}

func (r *ActorsPostgres) UpdateActorById(id int, new_actor models.Actor) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET name=$2, gender=$3, date_birth=$4, movie_id=$5 WHERE id=$1 RETURNING id", actorsTable)
	row := r.db.QueryRow(query, id, new_actor.Name, new_actor.Gender, new_actor.Date_birth, new_actor.Movie_id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActorsPostgres) DeleteActorById(id int) (int, error) {
	var deletedID int
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", actorsTable)
	err := r.db.QueryRow(query, id).Scan(&deletedID)
	if err != nil {
		return 0, err
	}
	return deletedID, nil
}
