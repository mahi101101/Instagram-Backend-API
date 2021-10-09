package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mahi101101/Instagram-Backend-API/routerm"
	"github.com/mahi101101/Instagram-Backend-API/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct{
	Session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (pc PostController) GetPost(w http.ResponseWriter, r *http.Request,p routerm.Params){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	po := models.Post{}

	if err := pc.Session.DB("mongo-golang").C("posts").FindId(oid).One(&po); err!=nil{
		w.WriteHeader(404)
		return
	}

	poj, err := json.Marshal(po)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",poj)
}

func (pc PostController) CreateNewPost(w http.ResponseWriter,r *http.Request, _ routerm.Params){
	po := models.Post{}

	json.NewDecoder(r.Body).Decode(&po)

	po.Id = bson.NewObjectId()
	po.TimeStamp = time.Now().String()
	pc.Session.DB("mongo-golang").C("posts").Insert(po)

	poj, err := json.Marshal(po)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",poj)
}

func (pc PostController) GetUserPosts(w http.ResponseWriter,r *http.Request, p routerm.Params){
	// userId := p.ByName("id")
	// if !bson.IsObjectIdHex(userId){
	// 	w.WriteHeader(404)
	// 	fmt.Println("3")
	// 	return
	// }


	// // po := models.Post{}

	// count,err := pc.Session.DB("mongo-golang").C("posts").Count()
	// if err!=nil{
	// 	w.WriteHeader(404)
	// 	fmt.Println("2")
	// 	return
	// }
	// fmt.Println(count)
	// //missing code
	// w.Header().Set("Content-type","application/json")
	// w.WriteHeader(http.StatusOK)

}
