package store

import (
	"goserver/model"

	"gorm.io/gorm"
)

func CreateUser(u *model.User, db *gorm.DB) error {
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(u *[]model.User, db *gorm.DB) error{
	err:= db.Find(u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(u *model.User, db *gorm.DB, id interface{}) error{
	err:= db.First(u,id).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(u *model.User, db *gorm.DB, id interface{}) error{
	err:= db.Delete(u,id).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(u *model.User, db *gorm.DB) error{
	err:= db.Save(u).Error
	if err != nil {
		return err
	}
	return nil
}




