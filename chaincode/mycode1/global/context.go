package global

import (
	ledgerapi "chaincode/ledger-api"
	"chaincode/mycode1/attribute"
	"chaincode/mycode1/user"
	"chaincode/mycode1/user_attribute"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetAttributeList() attribute.AttributeListInterface
	GetUserList() user.UserListInterface
	GetUserAttributeList() user_attribute.UserAttributeListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	attributeList     *attribute.AttributeList
	userList          *user.UserList
	userAttributeList *user_attribute.UserAttributeList
}

func (tc *TransactionContext) GetAttributeList() attribute.AttributeListInterface {
	if tc.attributeList == nil {
		tc.attributeList = NewAttrList(tc)
	}
	return tc.attributeList
}

func (tc *TransactionContext) GetUserList() user.UserListInterface {
	if tc.userList == nil {
		tc.userList = NewUserList(tc)
	}
	return tc.userList
}
func (tc *TransactionContext) GetUserAttributeList() user_attribute.UserAttributeListInterface {
	if tc.userAttributeList == nil {
		tc.userAttributeList = NewUserAttributeList(tc)
	}
	return tc.userAttributeList
}

// NewAttrList 创建一条属性链
func NewAttrList(ctx TransactionContextInterface) *attribute.AttributeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.studysdk.attrubutelist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return attribute.Deserialize(bytes, state.(*attribute.Attribute))
	}

	list := new(attribute.AttributeList)
	list.StateList = stateList
	return list
}

// NewUserList 创建一条用户链
func NewUserList(ctx TransactionContextInterface) *user.UserList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.studysdk.userlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return user.Deserialize(bytes, state.(*user.User))
	}

	list := new(user.UserList)
	list.StateList = stateList
	return list
}
func NewUserAttributeList(ctx TransactionContextInterface) *user_attribute.UserAttributeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.studysdk.user_attributelist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return user_attribute.Deserialize(bytes, state.(*user_attribute.UserAttribute))
	}

	list := new(user_attribute.UserAttributeList)
	list.StateList = stateList
	return list
}
