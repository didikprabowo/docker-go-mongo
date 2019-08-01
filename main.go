package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

var port = Env("APP_PORT")

type Post struct {
	Title string `bson:"title"`
}

func Env(name string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Get .env file")
	}

	env := os.Getenv(name)

	return env
}

func main() {

	http.HandleFunc("/", getPost)
	http.HandleFunc("/didik", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"APP_NAME": Env("APP_NAME"),
			"DB_HOST":  Env("DATABASE_HOST"),
			"DB_PORT":  Env("DATABASE_PORT"),
		}
		json.NewEncoder(w).Encode(&data)
	})

	http.ListenAndServe(":"+port, nil)
}

// ConnectMongo
func ConnectMongo() (*mgo.Database, error) {

	host := Env("DATABASE_HOST")
	port := Env("DATABASE_PORT")
	database := "blog"
	session, err := mgo.Dial(host + ":" + port)

	if err != nil {
		return nil, err
	}

	c := session.DB(database)
	return c, nil

}

// getPost
func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var posts []Post
	db, _ := ConnectMongo()
	db.C("posts").Find(bson.M{}).All(&posts)

	res := map[string]interface{}{
		"app_name": Env("APP_NAME"),
		"data":     posts,
	}

	json.NewEncoder(w).Encode(&res)
}
