package services

import (
	"hyperchain/api/models"
	"hyperchain/api/config"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var contract *gateway.Contract

func InitFabric() error {
	walletPath := config.AppConfig.WalletPath
	ccpPath := config.AppConfig.CCPPath
	identity := config.AppConfig.Identity

	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}

	if !wallet.Exists(identity) {
		return fmt.Errorf("%s identity not found in wallet", identity)
	}

	gw, err := gateway.Connect(
		gateway.WithConfigOption("discovery.asLocalhost", true),
		gateway.WithConfigFromPath(filepath.Clean(ccpPath)),
		gateway.WithIdentity(wallet, identity),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to gateway: %w", err)
	}

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return err
	}

	contract = network.GetContract(config.AppConfig.Chaincode)
	return nil
}

func AddRecord(record models.MedicalRecord) error {
	_, err := contract.SubmitTransaction("AddRecord", record.ID, record.Patient, record.Doctor, record.Diagnosis, record.Timestamp)
	return err
}

func QueryRecord(id string) (*models.MedicalRecord, error) {
	result, err := contract.EvaluateTransaction("QueryRecord", id)
	if err != nil {
		return nil, err
	}

	var record models.MedicalRecord
	if err := json.Unmarshal(result, &record); err != nil {
		return nil, err
	}
	return &record, nil
}
