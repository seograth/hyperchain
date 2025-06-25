package main

import (
    "log"
    "medical/contract"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
    cc, err := contractapi.NewChaincode(&contract.MediChainContract{})
    if err != nil {
        log.Fatalf("Error creating MediChain chaincode: %v", err)
    }

    if err := cc.Start(); err != nil {
        log.Fatalf("Error starting MediChain chaincode: %v", err)
    }
}