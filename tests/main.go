package main

import (
	"fmt"
	"lorm/models"
)

func main()  {
	user := models.NewUser()
	rs := []models.User{}
	user.Query().Where("id","=","20").Count("*").Sum("id").Get("id,name",rs)
	fmt.Println(len(rs))
	for _,v := range rs{
		fmt.Println(v.Id)
		fmt.Println(v.Name)
	}
}
