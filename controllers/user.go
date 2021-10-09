package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/mahi101101/Instagram-Backend-API/routerm"

	"github.com/mahi101101/Instagram-Backend-API/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s1 *mgo.Session) *UserController {
	return &UserController{s1}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p routerm.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.Session.DB("mongo-golang").C("users").FindId(oid).One(&u); err!= nil {
		w.WriteHeader(404)
		return
	}

	uj,err:= json.Marshal(u)
	if err!=nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ routerm.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	
	u.Id = bson.NewObjectId()
	pass := u.Password
	u.Password = string( encrypt([]byte(pass), "password"))

	uc.Session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}


