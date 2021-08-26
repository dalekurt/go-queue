# Queue 

Queueing service which uses Asnyq Go library to queue tasks and processing them asynchronously with workers using a Redis. 

## Getting started

1. Start a Redis server and asynqmon

```
docker-compose up -d
```

2. Go to [localhost:8080](http://localhost:8080/) to view queues

3. Start worker server

```console
make worker
```

4. Start generating tasks

```console
make client
```
