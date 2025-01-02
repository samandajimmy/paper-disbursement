# API Disbursement Service

Initialize project for Disbursement Service

## Prerequisites

1. Go 1.23.4
2. Git
3. Make sure you go install configured

   ```text
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go
   export GOBIN=$GOPATH/bin
   export PATH=$PATH:$GOPATH/bin
   ```

## Set-up

1. Configure Project

   ```sh
   # Run go prompt to download pkg dependencies
   go mod download
   ```

## Run Development

   ```sh
   # Run Service
   go build && ./paper-disbursement
   ```
