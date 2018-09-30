## To install all dependencies use
``glide install``

## Run redis in docker container

``docker run --name some-redis -d -p 6379:6379 redis``

### or local redis

## Run server
``go run server.go``
## Run client
``go run client.go``
### Then you can use console util gethashs
``go run gethashs.go``