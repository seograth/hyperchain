package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	CCPPath    string // Path to connection-org1.yaml
	WalletPath string // Path to wallet directory
	ChannelID  string // Fabric channel
	Chaincode  string // Chaincode name
	Identity   string // Wallet identity
}

var AppConfig *Config

func LoadConfig() {
	_ = godotenv.Load() // Load environment variables from .env file if it exists
	AppConfig = &Config{
		CCPPath:    getEnv("FABRIC_CCP_PATH", "../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml"),
		WalletPath: getEnv("FABRIC_WALLET_PATH", "wallet"),
		ChannelID:  getEnv("FABRIC_CHANNEL", "mychannel"),
		Chaincode:  getEnv("FABRIC_CHAINCODE", "medical"),
		Identity:   getEnv("FABRIC_IDENTITY", "admin"),
	}
	log.Println("Config loaded")
}

func getEnv(key string, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}
