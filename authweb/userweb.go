package authweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"nhms.com/auth/authrepo"
	"nhms.com/auth/authservice"
)

const (
	// ContentTypeKey for header content type
	ContentTypeKey = "Content-Type"
	// ContentType for response and request
	ContentType = "application/json"
	// Accept for response and request
	Accept = "application/json"
)

// GetUser an existing user the given user id.
func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	w.Header().Set(ContentTypeKey, ContentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := authservice.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// CreateUser a new user.
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var user authrepo.User
	err := decoder.Decode(&user)
	w.Header().Set(ContentTypeKey, ContentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = authservice.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// UpdateUser update an existing user with the given new user information.
func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("The rest has been called!")
	decoder := json.NewDecoder(r.Body)
	var updateuser authrepo.UpdateUser
	err := decoder.Decode(&updateuser)
	w.Header().Set(ContentTypeKey, ContentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = authservice.Update(&updateuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := authrepo.User{}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// DeleteUser delete an existing user.
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("Delete user has been called!"))
}

// GetUserByEmail get an existing user with the given email.
func GetUserByEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("Get user by email has been called!"))
}

// GetUserByPhoneNumber get user with the given phone number.
func GetUserByPhoneNumber(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("Get user by phone number has been called!"))
}
