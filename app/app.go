package app

import (
	"net/http"

	_"github.com/PeakActivity/go-todolist-challenge/app/lib"
	"gopkg.in/mgo.v2"
	_"gopkg.in/mgo.v2/bson"
	_"time"
	_"log"
	"time"
	"log"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)
const (
	Hosts = "ds237979.mlab.com:37979"
	Database = "todo"
	UserName = ""
	Password = ""
	httpSuccessCode = 200
	httpFailureCode = 401
)

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
	response := Result{}
	var allItems []Items
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	err = collection.Find(nil).All(&allItems)

	if err != nil {
		fmt.Println("Adding item err " + err.Error())
		response = Result{httpFailureCode, "Error occur in find Items"}
	} else {
		response = Result{httpSuccessCode, "successfully find Items"}

	}
	var out= map[string]interface{}{"response": response, "data": allItems}
	res.SendOK(out)

}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]
	response := Result{}
	var ItemById Items
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	err = collection.Find(bson.M{"_id":bson.ObjectIdHex(itemId)}).One(&ItemById)

	if err != nil {
		fmt.Println("Adding item err " + err.Error())
		response = Result{httpFailureCode,"Error occur in find Items"}
	}else {
		response = Result{httpSuccessCode,"successfully find Items"}

	}
	var out = map[string]interface{}{"response":response,"data":ItemById}
	res.SendOK(out)
}


func AddItems(w http.ResponseWriter, r *http.Request) {
	itemName := r.FormValue("name")

	response := Result{}
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	item := Items{bson.NewObjectId(),	itemName}
	err = collection.Insert(item)

	if err != nil {
		fmt.Println("Adding Item err " + err.Error())
		response = Result{httpFailureCode,err.Error()}
	}else {
		response = Result{httpSuccessCode,"successfully adding Items"}

	}

	resp, _ := json.Marshal(response)
	res.SendOK(resp)
}



func UpdateItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]
	itemName := r.FormValue("name")

	//var itemId = "ghjkl"
	response := Result{}
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	matchQuery := bson.M{"_id": bson.ObjectIdHex(itemId)}
	updateQuery := bson.M{"name":itemName}
	err = collection.Update(matchQuery, updateQuery)

	if err != nil {
		fmt.Println("updating item err " + err.Error())
		response = Result{httpFailureCode,err.Error()}
	}else {
		response = Result{httpSuccessCode,"successfully update Items"}

	}

	resp, _ := json.Marshal(response)
	res.SendOK(resp)
}


func DeleteItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]

	//var itemId = "ghjkl"
	response := Result{}
	res := lib.Response{ResponseWriter: w}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(Database).C("items")
	query := bson.M{"_id": bson.ObjectIdHex(itemId)}
	err = collection.Remove(query)

	if err != nil {
		fmt.Println("delete item err " + err.Error())
		response = Result{httpFailureCode,err.Error()}
	}else {
		response = Result{httpSuccessCode,"successfully delete Items"}

	}

	resp, _ := json.Marshal(response)
	res.SendOK(resp)
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