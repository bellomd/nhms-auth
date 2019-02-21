package authrepo

import (
	"fmt"
	"time"
)

// User is the model of the user.
type User struct {
	ID               int       `db:"id" json:"id"`
	FirstName        string    `db:"first_name" json:"firstName"`
	LastName         string    `db:"last_name" json:"lastName"`
	Email            string    `db:"email" json:"email"`
	Password         string    `db:"password" json:"password"`
	PhoneNumber      *string   `db:"phone_number" json:"phoneNumber"`
	Address          *string   `db:"address" json:"address"`
	CreatedBy        string    `db:"created_by" json:"createdBy"`
	LastModifiedBy   *string   `db:"last_modified_by" json:"lastModifiedBy"`
	CreatedDate      time.Time `db:"created_date" json:"createdDate"`
	LastModifiedDate time.Time `db:"last_modified_date" json:"lastModifiedDate"`
}

// UpdateUser is the user update dto that can be used
// to update a user with the given id.
type UpdateUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

// Create create a new user.
func (user *User) Create() (u *User, err error) {
	user.CreatedBy = CreatedBy
	user.CreatedDate = time.Now()
	tx := Db.MustBegin()
	//tx.NamedExec("INSERT INTO user(first_name, last_name, email, password, phone_number, address, created_by, last_modified_by, created_date,last_modified_date) VALUES (:first_name, :last_name, :email, :password, :phone_number, :address, :created_by, :last_modified_by, :created_date, :last_modified_date)", user)
	//tx.NamedExec("INSERT INTO t_user(first_name, last_name, email, password, phone_number, address, created_by, last_modified_by, created_date,last_modified_date) VALUES (:first_name, :last_name, :email, :password, :phone_number, :address, :created_by, :last_modified_by, :created_date, :last_modified_date)", &User{"Bello", "Muhammad", "bellomodigimba@gmail.com", "password", "phone", "address", "SYSTEM", "", time.Now(), time.Now()})
	//tx.MustExec("INSERT INTO t_user(first_name, last_name, email, password, phone_number, address, created_by, created_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.Address, user.CreatedBy, user.CreatedDate)
	rows, err := tx.NamedQuery("INSERT INTO t_user(first_name, last_name, email, password, phone_number, address, created_by, last_modified_by, created_date,last_modified_date) VALUES (:first_name, :last_name, :email, :password, :phone_number, :address, :created_by, :last_modified_by, :created_date, :last_modified_date) RETURNING id", user)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&user.ID)
	}
	tx.Commit()
	return
}

// GetUser get user with the given id.
func GetUser(ID int) (u *User, err error) {
	u = &User{}
	err = Db.Get(u, "SELECT * FROM t_user WHERE id=$1", ID)
	if err != nil {
		return nil, err
	}
	return
}

// Update update the given user.
func (user *User) Update() (u *User, err error) {
	fmt.Println("The repo has been called!")
	tx := Db.MustBegin()
	_, err = tx.NamedQuery("UPDATE t_user SET first_name:first_name, last_name:last_name, email:email, phone_number:phone_number, address:address, last_modified_by:last_modified_by, last_modified_date:last_modified_date", user)
	tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}

// DeleteUser delete user with the given id.
func DeleteUser(user *User) (u *User, err error) {
	return
}

// GetByEmail get user with the given email address.
func GetByEmail(email string) (u *User, err error) {
	return
}

// GetByPhoneNumber get user with the given phone number.
func GetByPhoneNumber(phoneNumber string) (u *User, err error) {
	return
}
