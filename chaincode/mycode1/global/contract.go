package global

import (
	"chaincode/mycode1/attribute"
	"chaincode/mycode1/user"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (mc *Contract) Instantiate() {
	fmt.Println("Contract Instantiated ")
}

// GetAllAttributes 获得账本中所有的属性记录
func (mc *Contract) GetAllAttributes(ctx TransactionContextInterface) ([]attribute.Attribute, error) {
	attributes, err := ctx.GetAttributeList().GetAllAttributes()
	if err != nil {
		return nil, err
	}
	return attributes, err

}

// AddAttribute 添加属性记录
func (mc *Contract) AddAttribute(ctx TransactionContextInterface, attr_name string) (*attribute.Attribute, error) {
	attr := attribute.Attribute(attr_name)
	err := ctx.GetAttributeList().AddAttribute(&attr)
	if err != nil {
		return nil, err
	}
	return &attr, err
}

// GetAttribute 查询属性记录
func (mc *Contract) GetAttribute(ctx TransactionContextInterface, attr_name string) (*attribute.Attribute, error) {
	attr, err := ctx.GetAttributeList().GetAttribute(attr_name)
	if err != nil {
		return nil, err
	}
	return attr, err
}

// DeleteAttribute 删除属性记录
func (mc *Contract) DeleteAttribute(ctx TransactionContextInterface, attr_name string) (*attribute.Attribute, error) {
	attr := attribute.Attribute(attr_name)
	err := ctx.GetAttributeList().DeleteAttribute(attr_name)
	if err != nil {
		return nil, err
	}
	return &attr, err
}

// GetUser 查询用户记录
func (mc *Contract) GetUser(ctx TransactionContextInterface, cerfiticate string) (*user.User, error) {
	user1, err := ctx.GetUserList().GetUser(cerfiticate)
	if err != nil {
		return nil, err
	}
	return user1, err
}

// AddUser 添加用户记录
func (mc *Contract) AddUser(ctx TransactionContextInterface, user_name string, certificate string) (*user.User, error) {
	newUser := user.User{UserName: user_name, Certificate: certificate, AttributeArray: []attribute.Attribute{}}
	err := ctx.GetUserList().AddUser(&newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, err
}

// UpdateUser 更新用户记录，只能修改用户名
func (mc *Contract) UpdateUser(ctx TransactionContextInterface, user_name string, certificate string) (*user.User, error) {
	updUser := user.User{UserName: user_name, Certificate: certificate, AttributeArray: []attribute.Attribute{}}

}
