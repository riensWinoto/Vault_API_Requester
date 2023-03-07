package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vault_api "github.com/hashicorp/vault/api"
)

var (
	vaultAddr  = os.Getenv("VAULT_ADDR")
	vaultToken = os.Getenv("VAULT_TOKEN")
	mountPath  = os.Args[1]
	secretPath = os.Args[2]
)

func initVaultClient() vault_api.Client {
	config := vault_api.Config{
		Address: vaultAddr,
	}
	client, err := vault_api.NewClient(&config)
	if err != nil {
		log.Fatalf("Unable to init vault client: %v", err)
	}
	client.SetToken(vaultToken)
	return *client
}

func fetchSecret(client *vault_api.Client, ctx context.Context, mountPath string, secretPath string) {
	secret, err := client.KVv2(mountPath).Get(ctx, secretPath)
	if err != nil {
		log.Fatalf("Unable to fetch secret: %v", err)
	}
	fmt.Println("Secret Fetched")
	for key, value := range secret.Data {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func main() {
	ctx := context.Background()
	client := initVaultClient()
	fetchSecret(&client, ctx, mountPath, secretPath)
}
