# Public RPC

<p align="left">
  <a href="https://github.com/web3toolz/public-rpc/actions"><img alt="GitHub Workflow Status (with event)" src="https://img.shields.io/github/actions/workflow/status/web3toolz/public-rpc/build_images_on_push.yaml"></a>
  &nbsp;
</p>

<hr/>
<h4>
<a target="_blank" href="https://public-rpc.web3toolz.com/" rel="dofollow"><strong>Website</strong></a>&nbsp;·
<a target="_blank" href="https://api-public-rpc.web3toolz.com/" rel="dofollow"><strong>Public API</strong></a>
</h4> 
<hr/>

### Features

* The most complete list of public RPC nodes for the top EVM and non-EVM compatible blockchains.
* Free Public API: Easy and open access for all developers. No API key required.

# Getting Started

## UI Component

**Install**

```shell
cd ui
yarn install
```

**Prepare configuration**

Copy .env.template to .env.local and edit it.

```shell
cp .env.template .env.local
```

**Run server in development mode**

Go to `https://localhost:3000` in your browser.

```shell
yarn dev
```

**Build static files**

They will be placed in `out` directory.

```shell
yarn build
```

## Backend component

**Install**

```shell
cd backend
go install 
```

**Prepare configuration**

Copy `.env.example` to `.env` to the same directory and edit it.

```shell
cp .env.example .env
```

**Run server in development mode**

```shell
go run cmd/cli/main.go run
```

**Request data from server**

```shell
curl -X GET http://localhost:8000/ | jq
```

**Build application**

```shell
go build -o public-rpc cmd/cli/main.go
```

## API Documentation

### Endpoints

* `GET "/"` - get public RPC nodes data

### Response format

The API returns data in JSON format. Here is an example of a successful response:

```json
[
  {
    "id": "75a38918-33f3-4277-a915-f3befb97283a",
    "http": "https://endpoints.omniatech.io/v1/eth/mainnet/public",
    "ws": "",
    "provider": "omniatech",
    "status": "active",
    "chainId": "1",
    "chain": "ethereum",
    "network": "mainnet",
    "addedAt": "2023-12-01T17:25:47.837Z",
    "checkedAt": "2023-12-01T17:25:47.837Z"
  },
  ...
]

```

## License

Distributed under the MIT License. See LICENSE for more information.