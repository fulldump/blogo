<img src="logo.png">

<p align="center">
<a href="https://travis-ci.org/fulldump/kip"><img src="https://travis-ci.org/fulldump/kip.svg?branch=master"></a>
<a href="https://goreportcard.com/report/fulldump/kip"><img src="http://goreportcard.com/badge/fulldump/kip"></a>
<a href="https://godoc.org/github.com/fulldump/kip"><img src="https://godoc.org/github.com/fulldump/kip?status.svg" alt="GoDoc"></a>
</p>


Kip is a Object wrapper for MongoDB.


<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [How to use](#how-to-use)
	- [Define](#define)
	- [Create DAO](#create-dao)
	- [CRUD: Create](#crud-create)
	- [CRUD: Retrieve](#crud-retrieve)
		- [FindOne](#findone)
		- [FindById](#findbyid)
		- [Find](#find)
	- [CRUD: Update](#crud-update)
	- [CRUD: Delete](#crud-delete)

<!-- /MarkdownTOC -->


# How to use

## Define

Basic usage:

```go
type User struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

kip.Define(&Collection{
	Name: "Users",
	OnCreate: func() interface{} {
		return &User{
			Id:   bson.NewObjectId(),
			Name: "default name",
		}
	},
})
```

Define indexes:

```go

kip.Define(&Collection{
	Name: "Users",
	OnCreate: func() interface{} {
		return &User{
			Id:   bson.NewObjectId(),
			Name: "default name",
		}
	},
}).EnsureIndex(mgo.Index{
    Key: []string{"email"},
    Unique: true,
    DropDups: true,
    Background: true, // See notes.
    Sparse: true,
})
```


## Create DAO

Definitions can be instantiated as many times as you want :)

```go
users := kip.Create("Users")
users.Database = NewDatabase("localhost", "demo")
```

## CRUD: Create

```go
john := users.Create()
```

## CRUD: Retrieve

Objects can be retrieved in three ways:

* FindOne
* FindById
* Find

### FindOne

Retrieve one item based on a query.

If there is no matching objects, nil is returned.

It will panic if an unexpected error happens.

```go
john := users.FindOne(bson.M{"name": "John"})
```

### FindById

Retrieve one item by `_id`.

It is a particular case of `FindOne` with the query `bson.M{"_id": <id> }`.


### Find

Retrieve a cursor...


## CRUD: Update

```go
TODO!
```

## CRUD: Delete

```go
john.Delete()
```
