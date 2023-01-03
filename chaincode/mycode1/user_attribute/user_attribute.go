package user_attribute

import (
	"chaincode/mycode1/attribute"
	"chaincode/mycode1/user"
	"encoding/json"
)

type UserAttribute struct {
	user           user.User             `json:"user"`
	attributeArray []attribute.Attribute `json:"attribute"`
}

func (ua *UserAttribute) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, ua)
}
func (ua *UserAttribute) MarshalJson() (data []byte, err error) {
	data, err = json.Marshal(ua)
	return data, err
}
func (ua *UserAttribute) Serialize() ([]byte, error) {
	return json.Marshal(ua)
}
func Deserialize(data []byte, ua *UserAttribute) error {
	return ua.UnmarshalJson(data)
}

//实现State接口
//两个函数Serialize()和GetSplitKey()

// GetSplitKey 返回StudentCard的主键字符串集
func (ua *UserAttribute) GetSplitKey() []string {
	return []string{ua.user.Certificate}
}
