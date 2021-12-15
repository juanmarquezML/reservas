package storage

import (
	"fmt"
	"learngo/api_only_go_task/cmd/utils"
	"learngo/api_only_go_task/pkg/model"
)

var users = make(map[string]*model.User)

type UserStorage interface {
	SaveUser(user *model.User) error
	GetUser(email string) (*model.User, error)
}
type userStorage struct{}

func NewUserStorage() *userStorage {
	return &userStorage{}
}

func (s *userStorage) SaveUser(user *model.User) error {
	/*
		if _, ok := users[user.Email]; !ok {
			users[user.Email] = user
		} else {
			log.Println(fmt.Sprintf("User %s already stored", user.Email))
			return fmt.Errorf("User %s already stored", user.Email)
		}

		return nil*/
	return nil
}

func (s *userStorage) GetUser(emaildir string) (*model.User, error) {
	sqlStatement := fmt.Sprintf("SELECT name, email FROM users where email = ?")
	//var selectValues []interface{}
	db := utils.InitDB()
	defer db.Close()
	userData := *&model.User{}
	//rows := db.QueryRow(sqlStatement, email)
	//rows, dbError := db.Query(sqlStatement, selectValues...)
	rows := db.QueryRow(sqlStatement, emaildir)
	rows.Err()
	//defer rows.Close()
	if rows.Err() != nil {
		panic(rows.Err())
	} else {
		//var name, email string
		//for rows.Next() {
		//rows.Scan(&name, &email)
		rows.Scan(&userData.Name, &userData.Email)
		//fmt.Println(email)
		//}

		//for rows.Next() {

		//}
	}
	return &userData, nil
	/*
		//err := row.Scan(&email)
		if err := rows.Scan(&email); err != nil {

		} else if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		}

		userData = model.User{
			Name:  "pepito",
			Email: email,
		}*/

}
