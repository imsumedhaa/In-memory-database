# In-Memory Database in Go
#### Introduction
An in-memory database is a database system that stores data primarily in the computer's main memory (RAM) instead of on traditional disk storage.

## How to run
Write **go run main.go inmemory** in the terminal and it will ask 'create,delete,update,get,show or exit' operations you want to perform and will ask the key and value depending on the operation. Then create a map and store it 
like map[key:value].

### Features

    Create → Add a key-value pair

    Get → Retrieve the value by key

    Update → Modify the value for an existing key

    Delete → Remove a key-value pair

    Show → View the entire store

    Exit → Quit the program

## Create a json file
Write **go run main.go filesystem --name db.json** in the terminal. If there is no file called db.json ,it will create json file called db.json. And same like inmemory, it will ask the operations.

