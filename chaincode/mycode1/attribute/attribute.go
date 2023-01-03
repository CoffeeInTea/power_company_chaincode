package attribute

import (
	"encoding/json"
	"fmt"
)

// Attribute 属性即一条字符串
type Attribute string

func (attr *Attribute) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, attr)
}

func (attr *Attribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(attr)
}

// Serialize 实现StateInterface
func (attr *Attribute) Serialize() ([]byte, error) {
	return json.Marshal(attr)
}

// GetSplitKey 实现StateInterface
func (attr *Attribute) GetSplitKey() []string {
	return []string{string(*attr)}
}
func Deserialize(bytes []byte, attr *Attribute) error {
	err := json.Unmarshal(bytes, attr)

	if err != nil {
		return fmt.Errorf("Error deserializing attribute. %s", err.Error())
	}

	return nil
}
