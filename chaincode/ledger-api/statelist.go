/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// StateListInterface functions that a state list
// should have
type StateListInterface interface {
	AddState(StateInterface) error
	GetState(string, StateInterface) error
	UpdateState(StateInterface) error
	GetAllStates(states []StateInterface) error
	DelState(state StateInterface) error
}

// StateList useful for managing putting data in and out
// of the ledger. Implementation of StateListInterface
type StateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
}

// AddState puts state into world state
func (sl *StateList) AddState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutState(key, data)
}

// GetState returns state from world state. Unmarshalls the JSON
// into passed state. Key is the split key value used in Add/Update
// joined using a colon
func (sl *StateList) GetState(key string, state StateInterface) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	data, err := sl.Ctx.GetStub().GetState(ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	return sl.Deserialize(data, state)
}

// UpdateState puts state into world state. Same as AddState but
// separate as semantically different
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}

// DelState 删除状态
func (sl *StateList) DelState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	return sl.Ctx.GetStub().DelState(key)
}

// GetAllStates 获得list中所有的记录
func (sl *StateList) GetAllStates(states []StateInterface) error {
	if len(states) != 0 {
		errors.New("states should be empty")
	}

	rs, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, []string{})
	if err != nil {
		fmt.Println(err)
	}
	defer func(rs shim.StateQueryIteratorInterface) {
		err := rs.Close()
		if err != nil {

		}
	}(rs)

	for rs.HasNext() {
		responseRange, err := rs.Next()
		if err != nil {
			fmt.Println(err)
		}
		state := new(StateInterface)
		err = json.Unmarshal(responseRange.Value, state)
		if err != nil {
			fmt.Println(err)
		}
		states = append(states, *state)
	}
	return nil
}
