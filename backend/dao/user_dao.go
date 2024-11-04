package dao

import "backend/model"

func (d *DAO) CreateUser(user model.User) {
	d.DB.Create(&user)
}
