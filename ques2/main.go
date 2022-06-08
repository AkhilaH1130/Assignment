// Q2. Create a middleware that log data in standard output. Please log
// following things: (25 Marks)
// • URL  For E.g.,  /home
// • Request Body E.g. {“name” : “Ana”}
// • Header data Hint: Check request object
// • Request Method E.g. GET
// • Hint: Check r *request object. We have methods already
// defined there to get relevant data.
// So, create an HTTP Server that accepts request on one endpoint and
// log the request data using middleware as mentioned above, and
// after logging data, route the request to the actual handler function
// where you can show or print some success message to the user.
// Please Note :- Send JSON body from the postman and log that using
// middleware.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

var c = content{}
var data []content

type content struct {
	Name string `json:"name"`
}

var Logger *logrus.Logger = logrus.New()

func logging(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Logger.Info("entering middleware")
		f.ServeHTTP(w, r)
	}
}

func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.SetFormatter(&logrus.JSONFormatter{})
		Logger.SetOutput(os.Stdout)
		Logger.Info("retrieving data")
		c_json, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		Logger.Info("data retrieved successfully")
		w.Write(c_json)
	})

}
func Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.SetFormatter(&logrus.JSONFormatter{})
		Logger.SetOutput(os.Stdout)
		Logger.Info("adding data to server")
		req, _ := ioutil.ReadAll(r.Body)

		err := json.Unmarshal(req, &c)

		if err != nil {
			fmt.Println(err)
		}
		data = append(data, c)
		Logger.Info("added successfully")
		w.Write(req)
	})

}

func main() {
	r := mux.NewRouter()

	r.Handle("/home",
		logging(
			Get())).Methods("GET")
	r.Handle("/home",
		logging(
			Create())).Methods("POST")
	http.ListenAndServe(":8080", r)
}
