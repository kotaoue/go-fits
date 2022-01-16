# go-fits
This script accesses the Fitness API

## Preparation
1. [Authenticating as a service account:Passing credentials manually](https://cloud.google.com/docs/authentication/production#cloud-console)

## Requirement
[google.golang.org/api/fitness/v1](https://pkg.go.dev/google.golang.org/api@v0.65.0/fitness/v1)
```
go get google.golang.org/api/fitness/v1
```

## Usage
```
export GOOGLE_PROJECT_ID="api-project-dummy"
export GOOGLE_APPLICATION_CREDENTIALS="./service-account-file.json"
```