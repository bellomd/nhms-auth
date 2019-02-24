package authservice

import (
	"time"

	"nhms.com/auth/autherror"
	"nhms.com/auth/authrepo"
)

// Get get user with the given id.
func Get(ID int) (u *authrepo.User, err error) {
	if ID <= 0 {
		return nil, autherror.New("ID cannot be less than zero(0).")
	}
	user, err := authrepo.GetUser(ID)
	if err != nil {
		return nil, autherror.New("User not found with the given id.")
	}
	return user, nil
}

// Create create a new user.
func Create(user *authrepo.User) (err error) {
	if user.Email == "" {
		return autherror.New("Email address cannot be empty.")
	}
	if user.Password == "" {
		return autherror.New("Password cannot be empty.")
	}
	user.Create()
	return
}

// Update update the given user.
func Update(user *authrepo.UpdateUserDto) (err error) {
	u, err := Get(user.ID)
	if err != nil {
		return
	}
	if isDirty(user, u) {
		err = u.Update()
	}
	return
}

// Delete delete user with the given id.
func Delete(ID int) (err error) {
	user, err := Get(ID)
	if err != nil {
		return
	}
	err = authrepo.DeleteUser(user)
	return
}

// GetByEmail get user with the given email address.
func GetByEmail(email string) (u *authrepo.User, err error) {
	if email == "" {
		return nil, autherror.New("Email cannot be empty.")
	}
	u, err = authrepo.GetByEmail(email)
	if err != nil {
		return nil, autherror.New("User not found with the given email address.")
	}
	return
}

// GetByPhoneNumber get user with the given phone number.
func GetByPhoneNumber(phoneNumber string) (u *authrepo.User, err error) {
	if phoneNumber == "" {
		return nil, autherror.New("Phone number cannot be empty.")
	}
	u, err = authrepo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, autherror.New("User not found with the given phone number.")
	}
	return
}

func isDirty(user *authrepo.UpdateUserDto, u *authrepo.User) bool {

	isDirty := false
	if user.FirstName != "" {
		u.FirstName = user.FirstName
		isDirty = true
	}
	if user.LastName != "" {
		u.LastName = user.LastName
		isDirty = true
	}
	if user.Email != "" {
		u.Email = user.Email
		isDirty = true
	}
	if user.PhoneNumber != "" {
		u.PhoneNumber = &user.PhoneNumber
		isDirty = true
	}
	if user.Address != "" {
		u.Address = &user.Address
		isDirty = true
	}
	u.LastModifiedBy = &u.Email
	u.LastModifiedDate = time.Now()

	return isDirty
}
