package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

func init() {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
}

type Movie struct {
	Title    string
	Year     int
	Director []string
	Cast     []string
	Genre    string
	notes    string
}

func insertData(data interface{}, messages chan bool) {
	session := MgoSession.Clone()
	clctn := session.DB("test").C("movies")
	err := clctn.Insert(&data)
	if err != nil {
		panic(err)
	}
	messages <- true
}
func main() {
	var data []interface{}
	url := "https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &data)
	//fmt.Printf("Results: %v\n", data)
	messages := make(chan bool, 4)
	fmt.Println(data[900])
	for i := 0; i < len(data); i++ {
		go insertData(data[i], messages)
		<-messages
	}

}
