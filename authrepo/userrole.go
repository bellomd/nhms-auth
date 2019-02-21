package authrepo

import "time"

// URole is the user's roles
type URole struct {
	ID               int       `json:"id"`
	Users            []User    `json:"users"`
	Roles            []Role    `json:"roles"`
	CreatedBy        string    `json:"createdBy"`
	LastModifiedBy   string    `json:"lastModifiedBy"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

// GetURole get user role with the given id.
func GetURole(ID int) (ur *URole, err error) {
	return
}

// Create create a new user role.
func (urole *URole) Create() (ur *URole, err error) {
	return
}

// Update update the given user role.
func (urole *URole) Update() (ur *URole, err error) {
	return
}

// DeleteURole delete user role with the given id.
func DeleteURole(ID int) (ur *URole, err error) {
	return
}
