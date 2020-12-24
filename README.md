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
export APP_AES_KEY=9a576e02a3beca3a64f2a1b5fa9061ea7d14716bf205d1b7c256ff7ef511a796
export ALPHA_VANTAGE_KEY=
```

### 4. Go to the git cloned folder and start the container

```
cd ~/src/github.com/ryuzaki01/go-ms
docker-compose up -d
```

### 5. Access with a browser (specify the IP address of the VM)

http://localhost:3000/symbol/TSLA

## Developer Notes (Challenges)

- What I think the hardest part while I was doing this project is, JSON operation, Parsing a huge json file would become a hard work
- Second is async function was pretty inconvenient, even thought I have created a helper for that it's still not good enough, problem is when we use async to fetch string, and then that string is used for another function param, the async function is returning interface instead of string, so we have to convert it to string
- http request operation a bit hard to be done if we don't use library
- found a weird thing in http request operation, if we do http request with method "POST" and we have an endpoint that have the trailing slash removed, ex: http://encrypt:3001/encypt the router still take it as a "GET" method
