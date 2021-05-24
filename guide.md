# Installing GO
> https://golang.org/doc/install

# Initialize go project
> go mod init place.holder/golist

# Compile and run
Execute at the project root directory
> go run .

# REST api
Following this tutorial
> https://tutorialedge.net/golang/creating-restful-api-with-golang/

# Testing the server
To test the server I mainly use curl commands
Examples:

curl -i -X GET localhost:1000
curl -i -X GET localhost:1000/lists
curl -i -X POST localhost:1000/create
curl -i -X GET localhost:10000/update -d '{"ID": 1, "items": [{"name": "chocolate", "check": true}]}'

# Commands
Currently command is independent of the verb. The existing commands are:

- lists: return all the existing lists
- create: create a new empty list
- delete: read the ID from the request body and delete the corresponding list.
- update: receive a list from the body and update the existing one with
	the same ID to look like it.
- store: Store the current lists to a file.
- load: Load the current lists from a file (the ones already present are
	discarded).
