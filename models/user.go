package models

import (
	"lorm"
	"reflect"
)


type User struct {
	*lorm.Model
	Name string `sql:"name"`
	Id int `sql:"id"`
}

func NewUser() *User{
	user := User{
		Model : &lorm.Model{
			TableName:  "users",
			PrimaryKey: "id",
			Connection: "default",
		},
	}
	user.Model.RealModel = user//
	user.Model.RealModelValue = reflect.ValueOf(&user).Elem()//
	return  &user
}
