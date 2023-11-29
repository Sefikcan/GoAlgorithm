## Installing dependencies

### Installing mongodb with docker

```
docker run --name mongo -p 27017:27017 -d mongo
```

### golang mongodb dependencies

```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
```

### Installing consul with docker

```
docker run -d \
  --name consul-container \
  -p 8300:8300 \
  -p 8301:8301 \
  -p 8301:8301/udp \
  -p 8500:8500 \
  -p 8600:8600 \
  -p 8600:8600/udp \
  bitnami/consul:latest 
```