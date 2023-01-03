package apply

import (
	"chaincode/mycode1/attribute"
	"chaincode/mycode1/user"
	"encoding/json"
)

type Apply struct {
	Applier     user.User             `json:"applier"`
	Attributes  []attribute.Attribute `json:"attributes"`
	State       string                `json:"state"`
	ApplyTime   string                `json:"apply_time"`
	ProcessTime string                `json:"process_time"`
}

func (a *Apply) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, a)
}
func (a *Apply) MarshalJson() (data []byte, err error) {
	data, err = json.Marshal(a)
	return data, err
}
func (a *Apply) Serialize() ([]byte, error) {
	return json.Marshal(a)
}
func Deserialize(data []byte, a *Apply) error {
	return a.UnmarshalJson(data)
}

//实现State接口
//两个函数Serialize()和GetSplitKey()

// GetSplitKey 返回StudentCard的主键字符串集
func (a *Apply) GetSplitKey() []string {
	return []string{a.user.Certificate}
}
