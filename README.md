# Datetime Client 

This repository implements an http datetime client that accepts the response of a datetime server

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/Datetime-client-RawanMostafa.git
   ```

2. Install the dependencies
    ```bash
    go mod download
    ```

## Usage

  You can configure our client using different ways:

### 1. Using flags

   ```bash
    go run cmd/main.go [flags]
   ```
#### Flags:
   - baseUrl=BASEURL 
   - port=PORT
   - endpoint=ENDPOINT

### 2. Using environment variables
This is used in case of no passed flags

   ```bash
    export VARNAME="my value"
    go run cmd/main.go 
   ```
#### Environment variables:
   - DATETIME_BASEURL 
   - DATETIME_PORT
   - DATETIME_ENDPOINT

### 2. Using the default configurations
Our application provides default configurations in case no flags are provided and environment variables aren't set

   ```bash
    go run cmd/main.go 
   ```
#### Default configs:
  ```go
    const defaltBaseUrl = "http://localhost"
    const defaultEndpoint = "/datetime"
    const defaultPort = "8083"
  ```
     
#### Extra Utilities
  - Check this function to get environment variables : [`decideConfigs`](https://github.com/codescalersinternships/Datetime-client-RawanMostafa/blob/9c3cc7ecf671057648e10d07c550c171b51747a6/cmd/main.go#L43-L76)
  - Check this function to read the returned response body : [`readBody`](https://github.com/codescalersinternships/Datetime-client-RawanMostafa/blob/9c3cc7ecf671057648e10d07c550c171b51747a6/cmd/main.go#L19-L28)
  - Check this function to get the flags : [`getFlags`](https://github.com/codescalersinternships/Datetime-client-RawanMostafa/blob/9c3cc7ecf671057648e10d07c550c171b51747a6/cmd/main.go#L29-L41)
   

