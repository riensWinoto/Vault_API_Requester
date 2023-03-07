# Vault_API_Requester
HashiCorp Vault API Requester

## How to use
Make sure set environment **VAULT_TOKEN** with your authorized token and **VAULT_ADDR** with your Vault IP address or DNS.

Install HashiCorp Vault API library
```
  go get github.com/hashicorp/vault/api
```

Run this command to run secret_fetcher.go
```
  go run secret_fetcher.go yourMountPath yourSecretPath
```
