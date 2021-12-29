- [1. Introduction <a name="intr"></a>](#1-introduction-)
- [2. Important links](#2-important-links)
- [3. Tech Stack and Architecture </a>](#3-tech-stack-and-architecture-a)
  - [3.1. Tech Stack](#31-tech-stack)
  - [3.2. Architecture](#32-architecture)
- [4. Database Schema </a>](#4-database-schema-a)
- [5. Usage Guide </a>](#5-usage-guide-a)
- [6. Development Notes </a>](#6-development-notes-a)
## 1. Introduction <a name="intr"></a>
Hosted on: heroku
In this project I aim to build simple ordering app and its simple service elements(authentication, ordering, order check)
My primary objective is creating extensible-scalable comment micro-service. My main application is just a tool to test/implement this microservice. 

I tried to use 
- Repository-Design
- Domain-Driven-Design
patterns while developing code.

I made load test using Jmeter and here the results:
    result

Used docker to create docker image and minikube to implement kubernetes-like service in local.

- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [minikube](https://minikube.sigs.k8s.io/docs/start/)

## 2. Important links
- [Live Application](heroku.com)
- [Jmeter Test Results](jmeter.com)
- [Swagger API Documentation](facebook.com)

## 3. Tech Stack and Architecture </a>
### 3.1. Tech Stack
  - Database
    - Postgresql
    - Redis (in-memory)
  - Swagger Documentation
    - https://github.com/swaggo/swag
### 3.2. Architecture
Clean Architecture

To demonstrate this I'll use this [tool](https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/) that auto-generate architecute.
## 4. Database Schema </a>
## 5. Usage Guide </a>
## 6. Development Notes </a>
- I'll not use any comments unless I think the code can unexpectedly behaves. Code needs to explain itself.