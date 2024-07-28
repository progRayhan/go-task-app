# To Do App
using Go Clean Architecture
Example that shows core principles of the Clean Architecture in Golang projects.

# Live Reload
[Live Reload](https://github.com/air-verse/air)
`go install` must be run in GOPATH. Otherwise, the binary will not be added to the path.
```bash
$ cd $GOPATH
$ go install github.com/cosmtrek/air@latest
```

# Setup
```bash
$ air init
```

# Database
We use MongoDB.
1. Start docker-compose 
2. For fresh database, we need to create the database manually (TODO)
```
$ mongosh
> use myroutine
> db.tasks.insert({title: "Hello, World!"})
```	

# Obtain Auth Token
1. See curl signup
2. See curl signin to obtain the token



## <a href="https://www.zhashkevych.com/clean-architecture">Blog Post</a>

## Rules of Clean Architecture by Uncle Bob:
- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Idependent of any external agency. In fact your business rules simply don’t know anything at all about the outside world. 

More on topic can be found <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">here</a>.

### Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing bookmarks (Simple clone of <a href="https://app.getpocket.com/">Pocket</a>).

#### Structure:
4 Domain layers:

- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## API:

### POST /auth/sign-up

Creates new user 

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```

##### Example Response: 
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /api/bookmarks

Creates new bookmark

##### Example Input: 
```
{
	"url": "https://github.com/mchayapol/go-task-app",
	"title": "Go Clean Architecture example"
} 
```

### GET /api/bookmarks

Returns all user bookmarks

##### Example Response: 
```
{
	"bookmarks": [
            {
                "id": "5da2d8aae9b63715ddfae856",
                "url": "https://github.com/mchayapol/go-task-app",
                "title": "Go Clean Architecture example"
            }
    ]
} 
```

### DELETE /api/bookmarks

Deletes bookmark by ID:

##### Example Input: 
```
{
	"id": "5da2d8aae9b63715ddfae856"
} 
```


## Requirements
- go 1.13
- docker & docker-compose

## Run Project for Production

Use ```make run``` to build and run docker containers with application itself and mongodb instance

## Run Project for Development
1. Requires an instance of MongoDB. See configuration in *config.yml*. The uri may need to change from
```
mongodb://mongodb:27017
```  
to
```
mongodb://127.0.0.1:27017
```
2. Run MongoDB instance
```
docker-compose -f mongo-compose.yml up -d
```

3. Run main.go  
```
go run cmd/api/main.go
```
## Run Automated Tests
1. Change to the directory containing the test
2. ```
go test -v handler_test.go handler.go register.go
```

or just run in each package, eg auth/delivery/http
```
go test -cover
```
or
Visualize coverage in VSCode by running `Go: Test Coverage in Current Package.

## Run Manual Tests
1. Sign up
```
curl -v -X POST -d '{"username":"mchayapol","password":"mchayapol"}' localhost:8000/auth/sign-up
```
2. Sign In
```
curl -v -X POST -d '{"username":"mchayapol","password":"mchayapol"}' localhost:8000/auth/sign-in
```
and obtain token
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE1NDQ2OTguNTc3OTc1LCJ1c2VyIjp7IklEIjoiNjY5YjVlMTQ5MGQ1ZTYxZTAxZmQxNmY1IiwiVXNlcm5hbWUiOiJtY2hheWFwb2wiLCJQYXNzd29yZCI6ImQwNWY0OTg5YjdkYzc1MTdhOWE2MTVkNDQ0ZmZjOGNmNDZhOTU5NTgifX0.-VQRoMqVb-KFJdyQKDgXhpxUh42fpCEE5SjkLKndzV0"}
```

4. Create a bookmark
```
curl -v -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIxNzMzMTIuMjI0MDk4LCJ1c2VyIjp7IklEIjoiNjZhNGY1ZjNhZDM5NmRjMzc5NzIzZmVlIiwiVXNlcm5hbWUiOiJtY2hheWFwb2wiLCJQYXNzd29yZCI6ImQwNWY0OTg5YjdkYzc1MTdhOWE2MTVkNDQ0ZmZjOGNmNDZhOTU5NTgifX0.kajEmWh56adozWTLbEucNP3w2C37VKBoJf0J2UnTJ9M" -d '{"url": "https://github.com/mchayapol/go-task-app","title": "Go Clean Architecture example"}' localhost:8000/api/bookmarks
```

5. Get all bookmarks
```
curl -v -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE1NDQ2OTguNTc3OTc1LCJ1c2VyIjp7IklEIjoiNjY5YjVlMTQ5MGQ1ZTYxZTAxZmQxNmY1IiwiVXNlcm5hbWUiOiJtY2hheWFwb2wiLCJQYXNzd29yZCI6ImQwNWY0OTg5YjdkYzc1MTdhOWE2MTVkNDQ0ZmZjOGNmNDZhOTU5NTgifX0.-VQRoMqVb-KFJdyQKDgXhpxUh42fpCEE5SjkLKndzV0" localhost:8000/api/bookmarks
```

# TODO
- Dev with Air (hot reload)
- Add Sentry (currently, logging is used.)