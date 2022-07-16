package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"
	"github.com/IchThoth/Go-MongoDB-REST-API/models"
)

type UserController struct {
	session  *mgo.Session
}

func NewUserController(s *mgo.Session)*UserController  {
	return &UserController{s}
}
func (uc *UserController)GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	id:= p.ByName("id")
	

	if !bson.ObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid:=bson.ObjectIdHex(id)

	u:= models.User{}

	uc.session.DB("mongo-golang").C("Users").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return 
	}
	uj, err := json.Marshal(u)
	if err!= nil {
		fmt.Fprintln(err)
	}
w.Header().Set("content-type", "application/json")
w.WriteHeader(http.StatusOK)
fmt.Printf(w , "%s/n" , uj)

}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id =bson.NewObjectId()
	uc.session.DB("mongo-golang").C("Users").Insert(u)
	uj, err := json.Marshal(u)
	if err!= nil {
		fmt.Fprintln(err)
	}
w.Header().Set("content-type", "application/json")
w.WriteHeader(http.StatusCreated)
fmt.Printf(w , "%s/n" , uj)
}
