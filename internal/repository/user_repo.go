package repository

import (
	"database/sql"
	"github.com/l-tting/fpl-app/internal/models"
)
//struct to store db connection
type UserRepository struct{
	DB  *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{DB:db}
}

func (r *UserRepository)Create(username,email,passwordHash string)(int,error){
	var id int
	err := r.DB.QueryRow(
		"insert into users(username,email,password_hash)values($1,$2,$3) returning id",
		username,email,passwordHash,
	).Scan(&id)

	if err != nil{
		return 0,err
	}
	return id,nil
}


func (r *UserRepository) GetByEmail(email string)(*models.User,error){
	u:= &models.User{}

	err := r.DB.QueryRow(
		"select id,username,password_haash,created_at from useers where email = $1",
		email,
	).Scan(&u.ID,&u.Username,&u.Email,&u.Password)

	if err != nil {
		return nil,err
	}
	return u,nil
}