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