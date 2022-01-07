- [1. Introduction <a name="intr"></a>](#1-introduction-)
- [2. Important links](#2-important-links)
- [3. Third-party libraries and Architecture </a>](#3-third-party-libraries-and-architecture-a)
  - [3.1. Tech Stack](#31-tech-stack)
  - [3.2. Architecture](#32-architecture)
- [4. Usage Guide </a>](#4-usage-guide-a)
- [5. Development Notes </a>](#5-development-notes-a)
## 1. Introduction <a name="intr"></a>

In this project I aim to build a microservice. 
This microservice should be extensible, scalable, cloud-native, containerized, well-tested comment micro-service. Project heavily inspired by [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) blog post by Uncle Bob. 
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

## 3. Third-party libraries and Architecture </a>
### 3.1. Tech Stack
  - Database
    - Postgresql
      - (Pgx Driver)[https://github.com/jackc/pgx]
      - (Squirrel -  Squirrel helps you build SQL queries from composable parts: )[https://github.com/Masterminds/squirrel]
  - Swagger Documentation
    - https://github.com/swaggo/swag
  - Logging
    - (Zerolog -  Zero Allocation JSON Logger) [https://github.com/rs/zerolog]
### 3.2. Architecture
Clean Architecture

To demonstrate this I'll use this [tool](https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/) that auto-generate architecute.
## 4. Usage Guide </a>
## 5. Development Notes </a>
- I'll not use any comments unless I think the code can unexpectedly behaves. Code needs to explain itself.
- It is assumed that default values are in yaml, and security-sensitive variables are defined in ENV.
- I made a choice about not getting involved to front-end side. 
- The JSON decoder does not report an error if values in the source do not correspond to values in the target. For example, it's not an error if the source contains the field "status", but the target does not.