package authweb

import (
	"encoding/json"
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
	decoder := json.NewDecoder(r.Body)
	var updateUserDto authrepo.UpdateUserDto
	err := decoder.Decode(&updateUserDto)
	w.Header().Set(ContentTypeKey, ContentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = authservice.Update(&updateUserDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteUser delete an existing user.
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("NotImplemented"))
}

// GetUserByEmail get an existing user with the given email.
func GetUserByEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("NotImplemented"))
}

// GetUserByPhoneNumber get user with the given phone number.
func GetUserByPhoneNumber(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set(ContentTypeKey, ContentType)
	w.Write([]byte("NotImplemented"))
}
