
# Task Management API

This is a simple Task Management API project built with Go. It allows you to create, update, retrieve, and delete tasks. The project includes four API endpoints and provides a Makefile for easy build and run commands.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Makefile Commands](#makefile-commands)

## Features

- **GET /tasks**: Retrieve a list of tasks.
- **POST /tasks**: Create a new task.
- **PUT /tasks/:id**: Update an existing task.
- **DELETE /tasks/:id**: Delete a task.

## Requirements

- Go 1.22 or higher
- Docker
- Make (for running the Makefile commands)

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

### 2. Build the Docker image:

```bash
make build
```

### 3. Run the application:

```bash
make run
```

## API Endpoints

### 1. GET `/tasks`

Retrieve a list of tasks.

#### Request:

```bash
curl -X GET http://localhost:8888/tasks?page=1&size=10
```

#### Response (200 OK):

```json
{
    "tasks": [
        {
            "id": "d7263666-a8b4-41bb-a5b5-e096b630489a",
            "name": "123",
            "status": 0
        }
    ],
    "size": 1,
    "page": 1
}
```

### 2. POST `/tasks`

Create a new task.

#### Request:

```bash
curl -X POST http://localhost:8888/tasks -H "Content-Type: application/json" -d '{
  "name": "New Task"
}'
```

#### Response (201 Created):
```bash
# No response body
```

### 3. PUT `/tasks/:id`

Update an existing task by ID.

#### Request:

```bash
curl -X PUT http://localhost:8888/tasks/task-1 -H "Content-Type: application/json" -d '{
  "name": "Updated Task",
  "status": 1
}'
```

#### Response (204 No Content):

```bash
# No response body
```

### 4. DELETE `/tasks/:id`

Delete a task by ID.

#### Request:

```bash
curl -X DELETE http://localhost:8888/tasks/task-1
```

#### Response (204 No Content):

```bash
# No response body
```

## Usage

1. **Build and run**: Use the Makefile to easily build and run the project in a Docker container.

2. **Interact with the API**: You can interact with the API via tools like `curl`, Postman, or directly through a frontend application.

## Makefile Commands

- **`make build`**: Builds the Docker image for the application.
- **`make run`**: Runs the application in a Docker container, exposing it on port 8080.
- **`make clean`**: Cleans up any Docker images and containers.
