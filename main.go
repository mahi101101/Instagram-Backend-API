package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/mahi101101/Instagram-Backend-API/controllers"
	"github.com/mahi101101/Instagram-Backend-API/routerm"

)

func main() {
	r := routerm.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)

	r.POST("/posts",pc.CreateNewPost)
	r.GET("/posts/:id",pc.GetPost)
	r.GET("/posts/users/:userid",pc.GetUserPosts)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
