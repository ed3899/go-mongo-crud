# Description
A simple CRUD and Restful app writen in Go and using MongoDB Atlas.
## Requirements
- Git <= 2.34.1 (Only for cloning this repo)
- Docker <=23.0.5
- Docker Compose <=3.
- [MongoDB Atlas Cluster](https://www.mongodb.com/docs/atlas/getting-started/)

### Optional
This extension allows you to populate your cluster with sample data. It allows you to run queries from your code editor as well.

- [MongoDB for VS Code](https://www.mongodb.com/products/vs-code)

## Run
Once you've created your MongoDB Cluster.

```
git clone https://github.com/ed3899/go-mongo-crud
touch .env
```

Populate the .env file with the following:

```
C1_DB_USERNAME=<YOUR_MONGO_DB_ATLAS_CLUSTER_USER>
C1_DB_PASSWORD=<YOUR_MONGO_DB_ATLAS_CLUSTER_PASSWORD>
C1_DB_CLUSTER=<YOUR_MONGO_DB_ATLAS_CLUSTER_NAME>
SERVING_PORT=8080
```

The following two you can get from the sample data in mentioned in the **Optional** section

```
C1_DB_AIRBNB=<YOUR_MONGO_DB_ATLAS_CLUSTER_DB_NAME>
C1_DB_AIRBNB_COLLEC_LISTINGS=<YOUR_MONGO_DB_ATLAS_CLUSTER_USER>
```

Run:
```
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