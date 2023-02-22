# Golang First Rest API

## What is it for
Simple sample of a REST API using Golang and Gin as http framework

## How to Run it
```
go run src/main.go
```

## Available endpoints

### GET

```
/albums
```
OBS: Returns a static slice of fake music albums

```
/albums/:id
```
OBS: Returns a specific object o a message informing that the object doesn't exist