Golang Microservice
---

## Description

Sample Microservice for retrieving stock data in Go Lang

## Services

Name               | Container Name  | Port (dev only)  | Description
------------------ | --------------- | ---------------- | ------------------
stock-ms           | stock           | 3000             | Stock Data Retriever
encrypt-ms         | encrypt         | 3001             | Data Encryptor

## Usage

### 1. Next preparation
  
Get Credentials for Alpha Vantage
https://www.alphavantage.co/support/#api-key

### 2. Start Docker Quickstart Terminal

### 3. Set environment variables

```
export SECRET=yourappsecret!@
export ALPHA_VANTAGE_KEY=
```

### 4. Go to the git cloned folder and start the container

```
cd ~/src/github.com/ryuzaki01/go-ms
docker-compose up -d
```

### 5. Access with a browser (specify the IP address of the VM)
