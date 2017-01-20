package main

import (
	mongo "mywebapp/libs/mongodb/db"
	mysql "mywebapp/libs/mysql/db"
)

func main() {
	mongodata := mongo.Get() //calling method of package  "mywebapp/libs/mongodb/db"
	sqldata := mysql.Get()   //calling method of package "mywebapp/libs/mysql/db"
	fmt.Println(mongodata)
	fmt.Println(sqldata)
}
