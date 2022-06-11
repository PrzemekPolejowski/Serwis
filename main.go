package main

import (
	"example.com/serw/models"
	"example.com/serw/routers"
)

var err error

func main() {
	//fmt.Println(controllers.ShowAllStudents)
	r := routers.SetupRouter()
	models.ConnectDatabase()
	r.Run()
}
