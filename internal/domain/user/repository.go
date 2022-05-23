package user

import (
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]User, error)
	FindOne(id int64) (*User, error)
}

const UsersCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

var settings = postgresql.ConnectionURL{
	Database: `eventdb`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `school22`,
}

func (r *repository) FindAll() ([]User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var users []User
	userColl := sess.Collection("users")
	err = userColl.Find().All(&users)
	if err != nil {
		log.Fatal("usersCol.FindAll: ", err)
	}
	return users, nil
}

func (r *repository) FindOne(id int64) (*User, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var user User
	userColl := sess.Collection("users")
	err = userColl.Find("id", id).One(&user)
	if err != nil {
		log.Fatal("userCol.FindOne: ", err)
	}

	return &user, nil
}
