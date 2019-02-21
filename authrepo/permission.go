package authrepo

import "time"

// Permission is the model of role and user permissions.
type Permission struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Value            string    `json:"value"`
	CreatedBy        string    `json:"createdBy"`
	LastModifiedBy   string    `json:"lastModifiedBy"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

// GetPermission get permission with the given id.
func GetPermission(ID int) (p *Permission, err error) {
	return
}

// Create create a new permission.
func (permission *Permission) Create() (p *Permission, err error) {
	return
}

// Update update the given permission.
func (permission *Permission) Update() (p *Permission, err error) {
	return
}

// DeletePermission delete permission with the given id.
func DeletePermission(ID int) (p *Permission, err error) {
	return
}

// GetByName get permission with the given name.
func GetByName(name string) (p *Permission, err error) {
	return
}

// GetByValue get permission with the given value.
func GetByValue(value string) (p *Permission, err error) {
	return
}
