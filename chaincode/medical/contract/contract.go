package contract

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MediChainContract struct {
	contractapi.Contract
}

func (m *MediChainContract) AddRecord(ctx contractapi.TransactionContextInterface, id, patient, doctor, diagnosis, timestamp string) error {
	record := MedicalRecord{
		ID:        id,
		Patient:   patient,
		Doctor:    doctor,
		Diagnosis: diagnosis,
		Timestamp: timestamp,
	}
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, recordJSON)
}

func (m *MediChainContract) QueryRecord(ctx contractapi.TransactionContextInterface, id string) (*MedicalRecord, error) {
	data, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("record %s does not exist", id)
	}

	var record MedicalRecord
	err = json.Unmarshal(data, &record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}
