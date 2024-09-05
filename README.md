# Task Optimizer (NOTE: THIS IS A WORK IN PROGRESS)
## PS: This repository may not be the most recent version of the project

Task Optimizer is a distributed microservice-based application designed to optimize task scheduling in real-time. It utilizes linear programming techniques to distribute tasks across multiple systems efficiently, minimizing time and resource costs. This system is ideal for use cases that require task delegation in dynamic, large-scale environments.


### Key Components:
- **Go**: Implements core microservices, ensuring high performance and concurrency.
- **Java**: Handles the linear programming algorithms and optimization engine.
- **Docker & Kubernetes**: For containerization and orchestration, enabling easy deployment and scaling.
- **PostgreSQL**: Used for storing task and resource data persistently.

## Tech Stack

- **Programming Languages**: Go, Java
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **APIs**: RESTful
- **Version Control**: Git/GitHub
- **Build Tools**: Gradle, Go Modules

## Prerequisites

Before you can run this project locally, ensure that you have the following installed:

- [Go](https://golang.org/doc/install) - v1.18 or above
- [Java JDK](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html) - v11 or above
- [Docker](https://docs.docker.com/get-docker/) - v20.10 or above
- [Kubernetes](https://kubernetes.io/docs/setup/) - v1.20 or above
- [Gradle](https://gradle.org/install/) - v8.0 or above
- [PostgreSQL](https://www.postgresql.org/download/) - v13 or above

## Installation

### Clone the Repository

```bash
git clone https://github.com/aniruddh-krovvidi/Task-Optimizer.git
cd Task-Optimizer
