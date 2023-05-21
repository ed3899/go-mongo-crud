# Description
A simple CRUD and Restful app writen in Go and using MongoDB Atlas.
## Requirements
- Git <= 2.34.1 (Only for cloning this repo)
- Docker <=23.0.5
- Docker Compose <=3.
## Run
```
git clone https://github.com/ed3899/go-mongo-crud
docker compose up -d
```

## API Testing
There's a file called `rest.json`. Import that in [ThunderClient](https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client) or [Postman](https://www.postman.com/) for testing endpoints.

Endpoints are described in `.request.[1].name` of the JSON.

# Development
## Requirements
- Go <=1.20.4
## How To Run
```
DB_USERNAME=YOUR_DB_USERNAME \
DB_PASSWORD=YOUR_DB_PASSWORD \
DB_NAME=YOUR_DB_NAME \
DB_CLUSTER=YOUR_DB_CLUSTER \
SERVING_PORT=YOUR_DESIRED_PORT
go run main.go
```