package user

import (
	"chaincode/mycode1/attribute"
	"encoding/json"
	"fmt"
)

type User struct {
	UserName       string                `json:"user_name"`       //用户名
	Certificate    string                `json:"certificate"`     //用户证书
	AttributeArray []attribute.Attribute `json:"attribute_array"` //属性集
}

func (user *User) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, user)
}

func (user *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(user)
}

// Serialize 实现StateInterface
func (user *User) Serialize() ([]byte, error) {
	return json.Marshal(user)
}

// GetSplitKey 实现StateInterface
func (user *User) GetSplitKey() []string {
	return []string{user.Certificate}
}
func Deserialize(bytes []byte, user *User) error {
	err := json.Unmarshal(bytes, user)

	if err != nil {
		return fmt.Errorf("Error deserializing User. %s", err.Error())
	}

	return nil
}
