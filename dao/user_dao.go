package dao

import models "go-postgres/model"

type UserDao interface {
	Select() ([]models.User, error)
	Insert(u models.User) error
}
