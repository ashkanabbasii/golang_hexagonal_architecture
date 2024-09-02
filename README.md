# Golang Hexagonal Architecture Example Microservice

This repository is an example implementation of a microservice built with **Golang**, following the principles of **Hexagonal Architecture** (also known as Ports and Adapters). The service is designed to be clean, maintainable, and scalable, with a focus on separation of concerns and high testability.

## Overview

This microservice demonstrates how to structure a Go application using hexagonal architecture, integrating with **PostgreSQL** as the primary database. The core business logic is decoupled from external dependencies, making it easier to adapt, test, and maintain.

### Key Features

- **Hexagonal Architecture**: The application is structured using the hexagonal architecture pattern, which separates the core business logic from external services like databases, APIs, etc.
- **PostgreSQL Integration**: The microservice uses PostgreSQL as its database, with all interactions managed through a repository layer.
- **Application Service Layer**: This implementation uses an application service layer to coordinate domain services, manage transactions, and handle the orchestration between different components of the application.

## Architecture

The project follows a well-defined directory structure based on hexagonal architecture principles:

- **/cmd**: Contains the application's entry point (main package).
- **/internal**: Houses the core business logic, domain services and models(entity), and application services.
- **/pkg**: Contains public and shared packages code.
- **/docs**: Contains swagger's related documents.

## Getting Started

To get started with this project, ensure you have Go and PostgreSQL installed on your machine.

### Prerequisites

- Go 1.23+ installed
- PostgreSQL installed and running

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ashkanabbasii/golang_hexagonal_architecture.git
   cd golang_hexagonal_architecture
