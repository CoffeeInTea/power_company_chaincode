package main

import (
	"chaincode/mycode1/global"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(global.Contract)
	contract.TransactionContextHandler = new(global.TransactionContext)
	contract.Name = "org.mycode.contract"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "MyCodeChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}
}
