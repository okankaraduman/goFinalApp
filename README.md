- [1. Introduction <a name="intr"></a>](#1-introduction-)
- [2. Important links](#2-important-links)
- [3. Third-party libraries and Architecture </a>](#3-third-party-libraries-and-architecture-a)
  - [3.1. Tech Stack](#31-tech-stack)
  - [3.2. Architecture and Project Layout](#32-architecture-and-project-layout)
- [4. Usage Guide </a>](#4-usage-guide-a)
- [5. Development Notes </a>](#5-development-notes-a)
- [TO-DO's </a>](#to-dos-a)
## 1. Introduction <a name="intr"></a>

In this project I aim to build a microservice. 
This microservice should be extensible, scalable, cloud-native, containerized, well-tested, monitorable. Projects architecture heavily inspired by [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) blog post by Uncle Bob. 
![Archi](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)


I tried to use 
- Repository-Design
- Domain-Driven-Design
patterns while developing code.

I made load test using Jmeter and here the results:
    result - will come

Used docker to create docker image and minikube to implement kubernetes-like service in local. 

- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [minikube](https://minikube.sigs.k8s.io/docs/start/)

## 2. Important links
- [Live Application](heroku.com) - will come
- [Jmeter Test Results](jmeter.com) - will come 
- [Swagger API Documentation](localhost:8080/swagger/index.html) - localhost:8080/swagger/index.html
- [Prometheus Metrics](localhost:8080/metrics) - localhost:8080/swagger/index.html

## 3. Third-party libraries and Architecture </a>
### 3.1. Tech Stack
  - Router
    - [Chi Router](https://github.com/go-chi/chi)
  - Database
    - Postgresql
      - [Pgx Driver](https://github.com/jackc/pgx)
      - [Squirrel -  Squirrel helps you build SQL queries from composable parts: ](https://github.com/Masterminds/squirrel)
  - Swagger Documentation
    - [Swaggo] (https://github.com/swaggo/swag)
  - Logging/Monitoring
    - [Zerolog -  Zero Allocation JSON Logger] (https://github.com/rs/zerolog)
    - [Prometheus Golang Client](https://github.com/prometheus/client_golang)
### 3.2. Architecture and Project Layout
```
project    
└───.github
│   └───workflows
│       │   ci.yml #Github Action Pipeline
└───cmd
│   └───app
│        │   main.go # Starting Point
└───config #reading from yaml&env
└───docs #Auto generated swagger docs
└───integration-test
└───internal
│   └───app
│   └───controller
│       └───http
│           └───v1
│               │  router.go
│               │  comment.go
│           └───v2
│               │  xx.go
│       └───amqp_rpc
│           │  router.go
│           │  comment.go
│   └───entity
│   └───usecase #Business logic.
│       │ interfaces.go 
│       │ comment.go 
│       │ mocks_test.go 
│       └───repo #A repository is an abstract storage (database) that business logic works with.
│           │  comment_postgres.go
│       └───webapi #It is an abstract web API that business logic works with. 
│           │  comment_api.go
│   └───migrations
│   └───pkg
│       └───httpserver
│       └───logger
│       └───postgres
│       └───rabbitmq
│ .env
│ .Dockerfile
│ .docker-compose

```

To demonstrate this I'll use this [tool](https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/) that auto-generate architecute.
## 4. Usage Guide </a>
## 5. Development Notes </a>
- I'll not use any comments unless I think the code can unexpectedly behaves. Code needs to explain itself.
- It is assumed that default values are in yaml, and security-sensitive variables are defined in ENV.
- I made a choice about not getting involved to front-end side. 
- The JSON decoder does not report an error if values in the source do not correspond to values in the target. For example, it's not an error if the source contains the field "status", but the target does not.
## TO-DO's </a>
- [ ] Translate the comment before adding/taking to/from database
- [ ] Correct implementation of yaml/env
- [ ] Write Usage Guide and test the application using just usage guide
- [ ] C4 Diagram generation