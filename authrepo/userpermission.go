package authrepo

import "time"

// UPermission is the users' permissions
type UPermission struct {
	ID               int          `json:"id"`
	Users            []User       `json:"users"`
	Permissions      []Permission `json:"permissions"`
	CreatedBy        string       `json:"createdBy"`
	LastModifiedBy   string       `json:"lastModifiedBy"`
	CreatedDate      time.Time    `json:"createdDate"`
	LastModifiedDate time.Time    `json:"lastModifiedDate"`
}

// GetUPermission get user permission with the given id.
func GetUPermission(ID int) (up *UPermission, err error) {
	return
}

// Create create a new user permission.
func (upermission *UPermission) Create() (up *UPermission, err error) {
	return
}

// Update update the given user permission.
func (upermission *UPermission) Update() (up *UPermission, err error) {
	return
}

// DeleteUPermission delete user permmission with the given id.
func DeleteUPermission(ID int) (up *UPermission, err error) {
	return
}
