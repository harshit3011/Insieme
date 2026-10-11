package repositories

import (
	"context"
	"log"

	"github.com/harshit3011/Insieme/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository{
	return &UserRepository{
		DB: db,
	}
}

func(r *UserRepository) CreateUser(ctx context.Context, user models.User) error{
	query:= `INSERT INTO users(name, username, email, passwordhash, profilepic, bio) VALUES($1, $2, $3, $4, $5, $6)`

	_,err:=r.DB.Exec(ctx,query, user.Name, user.Username, user.Email, user.PasswordHash, user.ProfilePic, user.Bio)
	if err!=nil{
		return err
	}
	log.Println("User successfully created!!")
	return nil
}