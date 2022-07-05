package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/IchThoth/Go-MongoDB-REST-API/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("")
	r.POST("")
	r.DELETE("")
}

func getSession() *mgo.Session {
	//connects to mongodb client//
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}
