package app

import (
	"net/http"

	"github.com/go-to-do/app/lib"
	"gopkg.in/mgo.v2"
	_"gopkg.in/mgo.v2/bson"
	_"time"
	_"log"
	"time"
	"log"
	_"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)
const (
	Hosts = "ds237979.mlab.com:37979"
	//Hosts = "localhost:27017"
	Database = "todo"
	UserName = "admin"
	Password = "admin"
	httpSuccessCode = 200
	httpFailureCode = 401
)
var response map[string]interface{}

var (
	mongoSession *mgo.Session
	err          error
)
type Items struct{
	ID					bson.ObjectId	 		`json:"_id" bson:"_id"`
	ItemName			string					`json:"item_name" bson:"item_name"`
}

type Result struct {
	code 	int
	message string
}


func init() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{Hosts},
		Timeout:  60 * time.Second,
		Database: Database,
		Username: UserName,
		Password: Password,
	}
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
}
//Message ...
type Message struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}

	m := Message{"Hello Get Your msg!"}
	res.SendOK(m)
}

// HelloWorld ...
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var allItems []Items
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	err = collection.Find(nil).All(&allItems)

	if err != nil {
		response = map[string]interface{}{"code":httpFailureCode,"message":"GetAllItems find err", "data": err.Error()}
	} else {
		response = map[string]interface{}{"code":httpSuccessCode,"message":"GetAllItems find success", "data": allItems}
	}
	fmt.Println("response",response)
	res.SendOK(response)

}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]
	var ItemById Items
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")

	err = collection.Find(bson.M{"_id":bson.ObjectIdHex(itemId)}).One(&ItemById)


	if err != nil {
		response = map[string]interface{}{"code":httpFailureCode,"message":"GetItemById find err", "data": err.Error()}
	} else {
		response = map[string]interface{}{"code":httpSuccessCode,"message":"GetItemById find success", "data": ItemById}
	}
	fmt.Println("response",response)
	res.SendOK(response)
}


func AddItems(w http.ResponseWriter, r *http.Request) {
	itemName := r.FormValue("name")
	fmt.Println("itemName",itemName)
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	item := Items{bson.NewObjectId(),	itemName}
	err = collection.Insert(item)

	if err != nil {
		response = map[string]interface{}{"code":httpFailureCode,"message":"AddItems find err", "data": err.Error()}
	} else {
		response = map[string]interface{}{"code":httpSuccessCode,"message":"AddItems find success", "data": item}
	}
	fmt.Println("response",response)
	res.SendOK(response)
}



func UpdateItems(w http.ResponseWriter, r *http.Request) {
	itemId := r.FormValue("id")
	itemName := r.FormValue("name")
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	matchQuery := bson.M{"_id": bson.ObjectIdHex(itemId)}
	updateQuery := bson.M{"item_name":itemName}
	fmt.Println("match",matchQuery,"updte",updateQuery)
	err = collection.Update(matchQuery, updateQuery)

	if err != nil {
		response = map[string]interface{}{"code":httpFailureCode,"message":"UpdateItems find err", "data": err.Error()}
	} else {
		response = map[string]interface{}{"code":httpSuccessCode,"message":"UpdateItems find success", "data": itemName}
	}
	fmt.Println("response",response)
	res.SendOK(response)
}


func DeleteItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	query := bson.M{"_id": bson.ObjectIdHex(itemId)}
	err = collection.Remove(query)

	if err != nil {
		response = map[string]interface{}{"code":httpFailureCode,"message":"DeleteItems find err", "data": err.Error()}
	} else {
		response = map[string]interface{}{"code":httpSuccessCode,"message":"DeleteItems find success", "data": ""}
	}
	fmt.Println("response",response)
	res.SendOK(response)
}

func HandleByMethod(w http.ResponseWriter, r *http.Request) {
	switch (r.Method) {
	case "GET" :
		break
	case "PUT" :
		break
	case "DELETE" :
		break

	}

}