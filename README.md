# In-Memory Database in Go
An in-memory database is a database system that stores data primarily in the computer's main memory (RAM) instead of on traditional disk storage.

### Run the application:
    go run main.go inmemory


It will ask 'create,delete,update,get,show or exit' operations you want to perform and will ask the key and value depending on the operation. Then create a map and store it 
like map[key:value].

### Features:

    Create → Add a key-value pair

    Get → Retrieve the value by key

    Update → Modify the value for an existing key

    Delete → Remove a key-value pair

    Show → View the entire store

    Exit → Quit the program

# File-Based Key-Value Store in Go
Key value database build in Golang that stores the data in a json file. Supports basic CRUD operations.

### Run the application:
    go run main.go filesystem --name db.json
 If db.json doesn't exist, it will be created automatically.
 All key-value pairs will be stored persistently in this file.  
# Postgres Key-Value Store in Go

This is a simple CLI-based key-value store application implemented with data persistence using a PostgreSQL database. It will create a table **'kvstore'** and store the key value in that table but not allow duplicate keys.

### Run the application:
    go run main.go postgres

### Require Dependencies:
    Go
    PostgreSQL
    github.com/lib/pq

## Create a Postgres comtainer using Docker
    docker run --name my-postgres-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=SecretPassword -e POSTGRES_DB=mydb -p 5431:5432 -d postgres
This command pulls the official PostgreSQL image from Docker Hub (if not already available) and starts a container.

**Note:** If PostgreSQL is already installed on your host machine and using the default port 5432, we map the container's 5432 port to 5431 on the host (-p 5431:5432) to
avoid conflicts.

## How to start Postgres
**Step 1:** Start the container(if not already running):

    docker start my-postgres-db
**Step 2:** Connect to the PostgreSQL database:

    psql -h localhost -p 5431 -U admin -d mydb

It will prompt for the password.

➡️ Enter the password you will be set, for example: SecretPassword


