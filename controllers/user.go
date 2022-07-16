package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IchThoth/Go-MongoDB-REST-API/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("Users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s/n", uj)

}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	uc.session.DB("mongo-golang").C("Users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s/n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("mongo-golang").C("Users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "/n")
}
