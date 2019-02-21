package authrepo

import "time"

// RPermission is the roles' permmissions
type RPermission struct {
	ID               int          `json:"id"`
	Roles            []Role       `json:"roles"`
	Permissions      []Permission `json:"permissions"`
	CreatedBy        string       `json:"createdBy"`
	LastModifiedBy   string       `json:"lastModifiedBy"`
	CreatedDate      time.Time    `json:"createdDate"`
	LastModifiedDate time.Time    `json:"lastModifiedDate"`
}

// GetRPermission get role permission with the given id.
func GetRPermission(ID int) (r *RPermission, err error) {
	return
}

// Create create a new role permission.
func (rpermission *RPermission) Create() (r *RPermission, err error) {
	return
}

// Update update the given role permission.
func (rpermission *RPermission) Update() (r *RPermission, err error) {
	return
}

// DeleteRPermission delete role permmission with the given id.
func DeleteRPermission(ID int) (r *RPermission, err error) {
	return
}
