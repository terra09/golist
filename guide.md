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
