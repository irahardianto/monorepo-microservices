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
---

## [Build The Project](https://github.com/irahardianto/monorepo-microservices#build-the-project)

### Run the go app

### Run using docker

---

## [Project Structure](https://github.com/irahardianto/monorepo-microservices#project-structure)

    /  
    |- bookings 
       |- httphandler
          |- httphandler.go
          |- resource.go
       |- grpcserver
       |- model
          |- model.go
       |- router
          |- router.go
       |- storage
          |- mongodb
          storage.go
    |- cmd
       |- bookings
          |- main.go
       |- bookingsgrpc
       |- movies
       |- moviesgrpc
       |- showtimes
       |- showtimesgrpc
       |- users
       |- usersgrpc
    |- config 
    |- k8s
       |- microservices
       |- ingress
    |- movies
    |- showtimes
    |- users
    |- utils
       |- logger.go
    |- vendor
    Dockerfile
    go.mod
    go.sum 

### bookings

### cmd

### config

### k8s

### utils

### vendor

### Dockerfile

### go.mod & go.sum

---

## [The Cmd Folder](https://github.com/irahardianto/monorepo-microservices#cmd)

---

 Implementation ToDos

- [x] Microservices implementation
- [x] Folder structuring
- [x] Env implementation using Viper
- [x] MongoDB using Atlas
- [x] Dockerfile to build services
- [ ] Custom logger to handle panic & other error
- [ ] Implement gRPC for the services
- [ ] Auth implementation with JWT
- [ ] Services JWT verification
- [ ] Services JWT usage
- [x] Kubernetes services deployment files
- [x] Kubernetes expose deployment internally via NodePort
- [ ] Kubernetes ingress configurations using Istio ingress controller
- [ ] CI pipeline using CircleCI, Travis?
- [ ] Explore Rook operator
- [ ] Federated Kubernetes

Readme ToDos

- [x] Intro
- [ ] Table of Content (Get Started)
- [ ] How to build the project
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