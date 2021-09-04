package repo

import (
	"log"
	"time"

	"github.com/jeffprestes/test-bbolt/conf"
	"github.com/jeffprestes/test-bbolt/model"
)

// SaveUser saves user's data into database
func SaveUser(group, email, name string, age int) (user model.User, err error) {
	user = model.User{
		Group:     group,
		Email:     email,
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
	}
	err = conf.UserDB.Save(&user)
	if err != nil {
		log.Println("[SaveUser] error saving user: ", err.Error(), user)
		return
	}
	return
}

func GetAllUsers() (users []model.User, err error) {
	err = nil
	err = conf.UserDB.All(&users)
	if err != nil {
		log.Println("[GetAllUsers] error getting all users: ", err.Error())
	}
	return
}
