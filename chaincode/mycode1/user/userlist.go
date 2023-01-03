package user

import (
	ledgerapi "chaincode/ledger-api"
	"errors"
)

type UserListInterface interface {
	AddUser(*User) error
	GetUser(string) (*User, error)
	UpdateUser(string) error
	DeleteUser(string) error
	GetAllUsers() ([]User, error)
}
type UserList struct {
	StateList ledgerapi.StateListInterface
}

func (ul *UserList) AddUser(user *User) error {
	return ul.StateList.AddState(user)
}

func (ul *UserList) GetUser(certificate string) (*User, error) {
	user := new(User)
	err := ul.StateList.GetState(certificate, user)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (ul *UserList) UpdateUser(certificate string) error {
	user, err := ul.GetUser(certificate)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("no such user exists, check the certificate")
	}
	return ul.StateList.UpdateState(user)
}
func (ul *UserList) DeleteUser(certificate string) error {
	user, err := ul.GetUser(certificate)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("no such user exists, check the certificate")
	}
	return ul.StateList.DelState(user)
}

func (ul *UserList) GetAllUsers() ([]User, error) {
	var users []User
	var states []ledgerapi.StateInterface
	err := ul.StateList.GetAllStates(states)
	if err != nil {
		return nil, err
	}
	for _, v := range states {
		user := v.(*User)
		users = append(users, *user)
	}
	return users, nil
}
