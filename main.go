package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Booking struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", returnAllBookings).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")
	myRouter.HandleFunc("/update-post/{id}", updateBooking).Methods("PUT")
	myRouter.HandleFunc("/delete/{id}", deleteBooking).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: new-booking POST")
	// get the body of the POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var booking Booking
	json.Unmarshal(reqBody, &booking)
	fmt.Printf("%+v\n", booking)
	db.Create(&booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBookings")
	bookings := []Booking{}
	db.Find(&bookings)
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		var booking Booking
		db.First(&booking, s)
		fmt.Println(booking)
		json.NewEncoder(w).Encode(booking)
	}
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var booking Booking
	json.Unmarshal(reqBody, &booking)
	fmt.Printf("%+v\n", booking)

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Update Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		// only update if id from path and json match
		if s == booking.ID {
			var dbBooking Booking
			db.Where("ID = ?", s).First(&dbBooking)
			dbBooking.User = booking.User
			dbBooking.Members = booking.Members
			db.Save(&dbBooking)
			fmt.Println(dbBooking)
			json.NewEncoder(w).Encode(dbBooking)
		}
	}
}

// delete a booking with the given id
func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Delete Booking No: ", key)
	s, err := strconv.Atoi(key)
	if err == nil {
		var booking Booking
		db.First(&booking, s)
		db.Delete(&booking)
	}
}

func main() {
	// Please define your username and password for MySQL.
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Estuary?charset=utf8&parseTime=True")
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
		db.AutoMigrate(&Booking{})
		handleRequests()
	}
}