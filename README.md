# Cinema Monorepo Microservices in Go

## Introduction

Hey! Welcome, The Cinema Monorepo Microservices project is an example of working implementation of Go for building multiple microservices using a single code repository. 

The project mainly wont be focusing on dealing with codes, rather than to deal with deployment using Kubernetes as the container orchestration platform. We'll look into how to deploy your containers to Kuberenetes, We'll also try to push further to implement cutting edge technologies such as Istio and use it as our Ingress, Statefulset using CSI, We'll also try to look into Rook which uses Ceph for our database storage provisioning. If you are interested in learning SOLID principles, dependency injection, unit-testing, and mocking, please refer to my other repo [service-pattern-go](https://github.com/irahardianto/service-pattern-go)

The Cinema backend is powered by 4 microservices written in Go, using MongoDB as its database.

- Movie Service: Provides information like movie ratings, title, etc.
- Show Times Service: Provides show times information.
- Booking Service: Provides booking information.
- Users Service: Provides movie suggestions for users by communicating with other services.

The project is based on the project written by [Manuel Morejón](https://github.com/mmorejon).

## Get Started
-  [Build The Project](https://github.com/irahardianto/monorepo-microservices#build-the-project)
-  [Project Structure](https://github.com/irahardianto/monorepo-microservices#project-structure)
-  [The Cmd Folder](https://github.com/irahardianto/monorepo-microservices#cmd)
-  [Todos](https://github.com/irahardianto/monorepo-microservices#todos)
---

## [Build The Project](https://github.com/irahardianto/monorepo-microservices#build-the-project)
First thing first, before we do anything else we want to make sure the application is working as expected. In order to build and run the application we can do it in two ways, the first one is to build your go app only, and the second one is to build your go app and embed it in docker then run it.

In both cases, before you run the app, you need to modify the mongodb connection string (atlasConnectionString) which is located in the config.yaml file to be able to connect to your mongodb atlas database, register for the free subscription [here](https://www.mongodb.com/cloud)

### Run the go app
Once you have change the connection string, you can build the go app for each respective services with the following command

```
go build -o {output-binary-name} github.com/irahardianto/monorepo-microservices/cmd/{service-to-build}

-------
example
-------

go build -o movies github.com/irahardianto/monorepo-microservices/cmd/movies
```

Then you can run each app by running the binary, note that we are binding to the same 8080 for all the services. 
```
./{output-binary-name} 

-------
example
-------

./movies
```
Then open your web browser and point to this url http://localhost:8080/{service} (e.g. http://localhost:8080/movies)

### Run using docker
We can also build our app and run in by using docker, the dockerfile is already included and you can see how the app is built by using docker.

We can trigger the build an run it by using following command
```
docker build -t irahardianto/cinema-movies:v1 --build-arg SERVICE_NAME=movies .
```

Then afterwards you can run it by using the following command 
```
docker run -p8080:8080 irahardianto/cinema-movies:v1    
```

We will discuss more on what are the arguments and flags involved in the docker section.

---

## [Project Structure](https://github.com/irahardianto/monorepo-microservices#project-structure)
    .
    ├── Dockerfile
    ├── bookings
    │   ├── httphandler
    │   │   ├── httphandler.go
    │   │   └── resource.go
    │   ├── grpcserver
    │   ├── model
    │   ├── router
    │   └── storage
    │       ├── mongodb
    │       │   └── storage.go
    │       └── storage.go
    ├── cmd
    │   ├── bookings
    │   │   └── main.go
    │   ├── bookingsgrpc
    │   ├── movies
    │   ├── moviessgrpc
    │   ├── showtimes
    │   ├── showtimesgrpc
    │   ├── users
    │   └── usersgrpc
    ├── config
    ├── go.mod
    ├── go.sum
    ├── k8s
        ├── helm
    │   ├── istio
    │   └── microservices
    ├── movies
    ├── showtimes
    ├── users
    └── vendor

While this is a monorepo, the services are loosely coupled and at any point in time in the future, the repo can be disintegrate easily into several repositories. The purpose of this is, it gives you options on how you want to design your organization, and build repositories and deployment mechanism around it.

### bookings, movies, showtimes, users
These are the folders consisting your microservice package, all of the modules required to run your microservices stored in each of its respective folder package.

* httphandler consist of http handler implementation, it is built by using chi router to maintain net/http compatibility, while still providing rich middleware.
* grpcserver the alternative implementation for the service, instead of using rest api, this one is implemented using gRPC. gRPC is a safe, secure, and very effecient way to communicate your microservices, it uses protobuff as its payload.
* model this is where the database model resides
* router router consist of routing mechanism for our httphandler
* storage storage is where we communicate with our database, in the root folder we define the base interface in storage.go and then we define the implementation to be then injected at runtime, in this case we implement mongodb. More on dependency injection and unit testing in my other repo [service-pattern-go](https://github.com/irahardianto/service-pattern-go)

### cmd
Cmd folder hosts your main.go file, with the folder, you can have 1 repo to be compiled as several binary.

### config
This is where the configuration lies, you can configure database connection string, database name, and application port. Currently we are using viper to read the yaml file, but in the future I may refactor it and uses Vault.

### k8s

### utils

### vendor

### Dockerfile
multi stage build

### go.mod & go.sum

---

## [The Cmd Folder](https://github.com/irahardianto/monorepo-microservices#cmd)

---

## [ToDos](https://github.com/irahardianto/monorepo-microservices#todos)

 Implementation Todos

- [x] Microservices implementation
- [x] Folder structuring
- [x] Env implementation using Viper
- [x] MongoDB using Atlas
- [x] Dockerfile to build services
- [ ] Custom logger to handle panic & other error
- [ ] Refactor Chi to Kataras/Muxie, if Muxie cost lower resources, benchmark!
- [ ] Refactor database connection to enable reconnection upon connection dropped
- [ ] Implement gRPC for the services
- [x] Auth implementation with JWT
- [ ] Services JWT verification
- [ ] Services JWT usage
- [x] Kubernetes services deployment files
- [x] Kubernetes expose deployment internally via NodePort
- [x] Kubernetes ingress configurations using Istio ingress controller
- [x] Kubernetes Istio egress configurations
- [x] Implement Liveness and Readiness probe for services
- [ ] TLS to secure ingress
- [x] Refactor Kubernetes deployment to Helm Chart
- [ ] CI pipeline using CircleCI, Travis?
- [ ] Explore Rook operator
- [ ] Federated Kubernetes

Readme Todos

- [x] Intro
- [ ] Table of Content (Get Started)
- [x] How to build the project
- [ ] Project structure
- [ ] Cmd
- [ ] Viper
- [ ] Go mod
- [ ] gRPC
- [ ] Auth with JWT
- [ ] Docker
- [ ] Multi-stage build
- [ ] Kubernetes
- [ ] CI Pipeline