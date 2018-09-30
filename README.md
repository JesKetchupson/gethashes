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
example ``go run gethashs.go 123456 2``
123456 - number to hash 
2 - how many times it will be changed and hased
