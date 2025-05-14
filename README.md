# In-Memory Database in Go
An in-memory database is a database system that stores data primarily in the computer's main memory (RAM) instead of on traditional disk storage.

### Run the application
    go run main.go inmemory


It will ask 'create,delete,update,get,show or exit' operations you want to perform and will ask the key and value depending on the operation. Then create a map and store it 
like map[key:value].

### Features

    Create → Add a key-value pair

    Get → Retrieve the value by key

    Update → Modify the value for an existing key

    Delete → Remove a key-value pair

    Show → View the entire store

    Exit → Quit the program

# File-Based Key-Value Store in Go
Key value database build in Golang that stores the data in a json file. Supports basic CRUD operations.

### Run the application
    go run main.go filesystem --name db.json
If fileis not there,it will create he file, then store the key value pairs in it.  
# Postgres Key-Value Store in Go

This is a simple CLI-based key-value store application implemented with data persistence using a PostgreSQL database. It will create a table **'kvstore'** and store the key value in that table but not allow duplicate keys.
    
### Need to setup
Create postgres connection (Can be done by Docker) with Postgres Username ,Postgres Password, Database Name, Port (default is 5432).

### Run the application
    go run main.go postgres

### **Require Dependencies**
    Go
    PostgreSQL
    github.com/lib/pq

