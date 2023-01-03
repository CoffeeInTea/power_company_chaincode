package attribute

import (
	"chaincode/ledger-api"
	"errors"
)

type AttributeListInterface interface {
	AddAttribute(*Attribute) error
	GetAttribute(string) (*Attribute, error)
	DeleteAttribute(string) error
	GetAllAttributes() ([]Attribute, error)
}

type AttributeList struct {
	StateList ledgerapi.StateListInterface
}

// AddAttribute 添加属性记录
func (attrl *AttributeList) AddAttribute(attribute *Attribute) error {
	return attrl.StateList.AddState(attribute)
}

// GetAttribute 根据属性名获得属性记录
func (attrl *AttributeList) GetAttribute(attr_name string) (*Attribute, error) {
	attr := new(Attribute)
	err := attrl.StateList.GetState(attr_name, attr)
	if err != nil {
		return nil, err
	}
	return attr, err
}

// DeleteAttribute 删除属性
func (attrl *AttributeList) DeleteAttribute(attr_name string) error {
	attr, err := attrl.GetAttribute(attr_name)
	if err != nil {
		return err
	}
	if attr == nil {
		return errors.New("no such attribute exists")
	}
	err = attrl.StateList.DelState(attr)
	return err
}

// GetAllAttributes 获得账本中所有属性记录
func (attrl *AttributeList) GetAllAttributes() ([]Attribute, error) {
	var attributes []Attribute
	var states []ledgerapi.StateInterface
	err := attrl.StateList.GetAllStates(states)
	if err != nil {
		return nil, err
	}
	for _, v := range states {
		attr := v.(*Attribute)
		attributes = append(attributes, *attr)
	}
	return attributes, nil
}
