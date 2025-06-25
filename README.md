# 🩺 Hyperchain – Medical Records POC on Hyperledger Fabric (Go)

Hyperchain is a Proof-of-Concept blockchain-based medical record system built on **Hyperledger Fabric** using **Go chaincode** and a **Go backend API**. It is designed to provide a clean architecture, follow enterprise standards, and support easy onboarding for junior developers.

## Project Structure

```bash
hyperchain/
├── api/                          # Backend REST API written in Go
│   ├── config/                   # Environment configuration logic
│   ├── controllers/             # HTTP handlers
│   ├── routes/                  # Route definitions
│   ├── services/                # Business logic, incl. Fabric service
│   ├── utils/                   # Helper utilities
│   └── cmd/                  # API entrypoint
│
├── chaincode/                   # Chaincode (smart contract) folder
│   └── medical/                 # 'medical' chaincode project
│       ├── contract/            # Contract logic (handlers, models)
│       │   ├── contract.go      # Fabric smart contract implementation
│       │   └── models.go        # Data model structs
│       ├── main.go              # Chaincode main entrypoint
│       └── go.mod               # Chaincode module definition
│
└── README.md                    # You're here
```

## Prerequisites

Make sure the following are installed on your system:

-   Go (>= 1.20): https://go.dev/dl/
-   Node.js + npm (for Fabric tooling): https://nodejs.org
-   Docker + Docker Compose: https://docs.docker.com/get-docker/
-   Hyperledger Fabric Samples: https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html
-   WSL2 (if on Windows): https://learn.microsoft.com/en-us/windows/wsl/

Note: Keep fabric-samples outside the project folder (e.g., ~/fabric-samples).

## Environment setup

Create a .env file inside the api/ directory with the following content:

```bash
env
FABRIC_CCP_PATH=/home/<your-user>/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml
FABRIC_WALLET_PATH=/home/<your-user>/hyperchain/api/wallet
FABRIC_CHANNEL=mychannel
FABRIC_CHAINCODE=medical
FABRIC_IDENTITY=admin
```

Adjust file paths according to your WSL or Linux home structure.

## Fabric Network Setup

```bash
bash
cd ~/fabric-samples/test-network
```

Then bring up the network and deploy the chaincode:

```bash
./network.sh down
./network.sh up createChannel -ca
./network.sh deployCC -ccn medical -ccp ~/hyperchain/chaincode/medical -ccl go
```
Verify that deployment is successful.

## Run the API

From the `api/` folder:

```bash
cd ~/hyperchain/api
go run main.go
```
Your REST API should be running on `http://localhost:8080`.

## API Endpoints

| Endpoint               | Method | Description                  |
|---------------------|--------|------------------------------|
| `/api/record`          | POST   | Add a new medical record     |
| `/api/record/{id}`     | GET    | Query a medical record by ID |


## Chaincode Overview

Chaincode is located in `chaincode/medical`:

- `main.go`: Boots up the Fabric chaincode runtime
- `contract/contract.go`: Implements business logic (e.g., `AddRecord`)
- `contract/models.go`: Contains the `MedicalRecord` struct

The chaincode is a standalone Go module and should be built separately if needed using:

```bash
cd chaincode/medical
go build
```

## Identity & Wallet

Before calling the chaincode, enroll and import a user identity (`admin`) using Fabric CA tools or Fabric gateway SDK.

The identity must be stored in the wallet path defined in `.env`:

```bash
/api/wallet/admin.id
/api/wallet/admin.cert
/api/wallet/admin.key
```

## Developer Guide

- Changes to chaincode require re-deploying using `deployCC`
- Use Postman, curl, or Swagger UI to test API calls
- Logs from Fabric containers help debug chaincode execution
- Use `go mod tidy` and consistent import paths in Go

## Useful Commands

Build chaincode locally:

```bash
cd chaincode/medical
go build
```
Query committed chaincodes:

```bash
peer lifecycle chaincode querycommitted --channelID mychannel --name medical
```
Inspect wallet contents:

```bash
ls api/wallet
```
##  Common Issues

- **Chaincode "not found in registry"**: Check hash mismatch, rebuild, and re-deploy
- **Import cycle errors**: Avoid circular imports and organize packages clearly
- **Function does not exist**: Ensure the chaincode method is public and correctly named
- **Core config not found**: Ensure you're in the correct WSL directory when running Fabric CLI

##  Roadmap

- [ ] Chaincode methods for update and delete
- [ ] Query medical records by patient or doctor
- [ ] Add authentication to API (JWT or Fabric CA login)
- [ ] Dockerize Go API for dev/test deployments
- [ ] Frontend dashboard for medical record management

##  License

MIT License — © 2025 Hyperchain Contributors

