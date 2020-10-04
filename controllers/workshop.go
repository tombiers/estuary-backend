package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/models"
	"github.com/tombiers/estuary-backend/services"
)

// CreateNewWorkshop create a new workshop
func CreateNewWorkshop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateNewWorkshop POST")
	// get the body of the POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var workshop models.Workshop
	json.Unmarshal(reqBody, &workshop)
	services.CreateWorkshop(workshop)
}

// AllWorkshops returns all known workshops
func AllWorkshops(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AllWorkshops")
	workshops := services.GetAllWorkshops()
	json.NewEncoder(w).Encode(workshops)
}

// SingleWorkshop returns a single workshop
func SingleWorkshop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: Workshop No: ", key)

	var workshop = services.GetWorkshopByUUID(key)
	fmt.Println(workshop)
	json.NewEncoder(w).Encode(workshop)
}

// UpdateWorkshop updates the workshop with the given uuid
// reuqested uuid and uuid in the update data have to match
func UpdateWorkshop(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var workshop models.Workshop
	json.Unmarshal(reqBody, &workshop)

	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: Update Workshop No: ", key)

	// only update if id from path and json match
	if key == workshop.UUID {
		var newWorkshop = services.UpdateWorkshop(key, workshop)
		json.NewEncoder(w).Encode(newWorkshop)
	}

}

// DeleteWorkshop deletes the workshop with the given uui
func DeleteWorkshop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: Delete Workshop No: ", key)
	services.DeleteWorkshop(key)
}