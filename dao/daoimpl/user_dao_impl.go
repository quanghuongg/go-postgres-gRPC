package daoimpl

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"go-postgres/dao"
	models "go-postgres/model"
)

type UserDaoImpl struct {
	Db *sql.DB
}

func (u UserDaoImpl) Select() ([]models.User, error) {
	users := make([]models.User, 0)
	rows, err := u.Db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Name, &user.Id)
		if err != nil {
			break
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u UserDaoImpl) Insert(user models.User) error {
	insertStatement := `INSERT INTO users(name) VALUES($1)`
	_, err := u.Db.Exec(insertStatement, user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User added: ", user)
	return nil
}

func NewUseDao(db *sql.DB) (userDao dao.UserDao) {
	return &UserDaoImpl{
		Db: db,
	}
}
