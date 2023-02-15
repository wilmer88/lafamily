package models

import (
	
	"gorm.io/gorm"
)

type Fammember struct {
	gorm.Model
	Id   int
	FirstName  string
	Happiness int
	UrlStr string
}

//create a user
func CreateUser(db *gorm.DB, User *Fammember) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]Fammember) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *Fammember, Id int) (err error) {
	err = db.Where("Id = ?", Id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *Fammember) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *Fammember, Id int) (err error) {
	db.Where("Id = ?", Id).Delete(User)
	return nil
}