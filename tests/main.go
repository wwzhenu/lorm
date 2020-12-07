package main

import (
	"fmt"
	"lorm/models"
)

func main()  {
	user := models.NewUser()
	rs := ""
	//user.Query().Where("id","<","10").Count("*").Sum("id").Limit(2).Get("id,name",&rs)
	//user.Query().Where("id","=","10").Count("*").Sum("id").First("id,name",&rs)
	user.Query().Where("id","=","10").Count("*").Sum("id").Value("name",&rs)
	fmt.Println(rs)
}
