/**
 * 
 * Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
 * 
 */
package main

import (
	"fmt"
	"example.com/gitbugtestgo1/lib/chaincode"
	"example.com/gitbugtestgo1/lib/util"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	util.ChaincodeName = "gitbugtestgo1"
	err := shim.Start(new(chaincode.ChainCode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
