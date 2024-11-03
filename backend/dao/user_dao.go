package dao

import "Score_System/model"

func (d *DAO) CreateUser(user model.User) {
	d.DB.Create(&user)
}
