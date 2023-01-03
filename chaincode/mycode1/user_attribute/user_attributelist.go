package user_attribute

import (
	ledgerapi "chaincode/ledger-api"
	"chaincode/mycode1/user"
)

type UserAttributeListInterface interface {
	AddUserAttribute(*UserAttribute) error
	GetUserAttribute(string) (*UserAttribute, error)
	DelUserAttribute(string) error
	GetAllUserAttribute() ([]UserAttribute, error)
}
type UserAttributeList struct {
	StateList ledgerapi.StateListInterface
}

func (ual *UserAttributeList) AddUserAttribute(au *UserAttribute) error {
	return ual.StateList.AddState(au)
}
func (ual *UserAttributeList) GetUserAttribute(certificate string) (*UserAttribute, error) {
	userAttribute := new(UserAttribute)
	err := ual.StateList.GetState(certificate, userAttribute)
	if err != nil {
		return nil, err
	}
	return userAttribute, err
}
func (ual *UserAttributeList) DelUserAttribute(certificate string) error {
	userAttribute := UserAttribute{user: user.User{Certificate: certificate}}
	err := ual.StateList.DelState(&userAttribute)
	return err
}
func (ual *UserAttributeList) GetAllUserAttribute() ([]UserAttribute, error) {
	var userAttributeArray []UserAttribute
	var states []ledgerapi.StateInterface
	err := ual.StateList.GetAllStates(states)
	if err != nil {
		return nil, err
	}
	for _, v := range states {
		userAttribute := v.(*UserAttribute)
		userAttributeArray = append(userAttributeArray, *userAttribute)
	}
	return userAttributeArray, nil
}
