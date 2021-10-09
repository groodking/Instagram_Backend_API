package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User - Our struct for all Users
type User struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Post struct {
	Id               string `json:"Id"`
	Caption          string `json:"Caption"`
	Image_URL        string `json:"Image URL"`
	Posted_Timestamp string `json:"Timestamp"`
}

var Users []User
var Posts []Post

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, User := range Users {
		if User.Id == key {
			json.NewEncoder(w).Encode(User)
		}
	}
}

func returnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, Post := range Posts {
		if Post.Id == key {
			json.NewEncoder(w).Encode(Post)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new User struct
	// append this to our Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var User User
	json.Unmarshal(reqBody, &User)
	// update our global Users array to include
	// our new User
	Users = append(Users, User)

	json.NewEncoder(w).Encode(User)
}
func createNewPost(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new User struct
	// append this to our Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Post Post
	json.Unmarshal(reqBody, &Post)
	// update our global Users array to include
	// our new User
	Posts = append(Posts, Post)
	json.NewEncoder(w).Encode(Post)
}

func returnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	json.NewEncoder(w).Encode(Posts)

	vars := mux.Vars(r)
	key := vars["id"]
	key2 := vars["name"]

	for _, Post := range Posts {
		for _, User := range Users {
			if Post.Id == key && User.Name == key2 {
				json.NewEncoder(w).Encode(Post)
			}
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", returnSingleUser)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func handleRequests1() {
	myRouter1 := mux.NewRouter().StrictSlash(true)
	myRouter1.HandleFunc("/", homePage)
	myRouter1.HandleFunc("/posts", createNewPost).Methods("POST")
	myRouter1.HandleFunc("/posts/{id}", returnSinglePost)
	myRouter1.HandleFunc("/posts/users/{id}", returnAllPosts)
	log.Fatal(http.ListenAndServe(":10000", myRouter1))
}
func main() {
	Users = []User{
		User{Id: "1", Name: "Saransh", Email: "saransh.123@gmail.com", Password: "1234"},
		User{Id: "2", Name: "Ayush", Email: "Ayushmehra123@yahoo.com", Password: "1236"},
	}
	handleRequests()

	Posts = []Post{
		Post{Id: "1", Caption: "Hello", Image_URL: "logo.png", Posted_Timestamp: "Timestamp"},
		Post{Id: "2", Caption: "Hello 2", Image_URL: "logo1.png", Posted_Timestamp: "Timestamp"},
	}
	handleRequests1()
}
