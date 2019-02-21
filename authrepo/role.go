package authrepo

import "time"

// Role is the model of user roles
type Role struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CreatedBy        string    `json:"createdBy"`
	LastModifiedBy   string    `json:"lastModifiedBy"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

// GetRole get role with the given id.
func GetRole(ID int) (r *Role, err error) {
	return
}

// Create create a new role.
func (role *Role) Create() (r *Role, err error) {
	return
}

// Update update the given role.
func (role *Role) Update() (r *Role, err error) {
	return
}

// DeleteRole delete role with the given id.
func DeleteRole(ID int) (r *Role, err error) {
	return
}

// GetRoleByName get role with the given name.
func GetRoleByName(name string) (r *Role, err error) {
	return
}
