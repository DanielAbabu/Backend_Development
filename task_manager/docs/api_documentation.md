# Task Management API Documentation

This document provides an overview of the Task Management API, including available endpoints and example requests and responses.

## Base URL

```
https://api.example.com
```

## Endpoints

### Create Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Description**: Create a new task.
- **Request Body**:
  ```json
  {
    "title": "Task title",
    "description": "Task description",
    "status": "Task status"
  }
  ```
- **Response**:
  ```json
  {
    "id": "Task ID",
    "title": "Task title",
    "description": "Task description",
    "status": "Task status"
  }
  ```

### Get Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieve a list of all tasks.
- **Response**:
  ```json
  [
    {
      "id": "Task ID",
      "title": "Task title",
      "description": "Task description",
      "status": "Task status"
    }
  ]
  ```

### Get Task by ID

- **URL**: `/tasks/{id}`
- **Method**: `GET`
- **Description**: Retrieve a task by its ID.
- **Response**:
  ```json
  {
    "id": "Task ID",
    "title": "Task title",
    "description": "Task description",
    "status": "Task status"
  }
  ```

### Update Task

- **URL**: `/tasks/{id}`
- **Method**: `PUT`
- **Description**: Update a task by its ID.
- **Request Body**:
  ```json
  {
    "title": "Updated title",
    "description": "Updated description",
    "status": "Updated status"
  }
  ```
- **Response**:
  ```json
  {
    "id": "Task ID",
    "title": "Updated title",
    "description": "Updated description",
    "status": "Updated status"
  }
  ```

### Delete Task

- **URL**: `/tasks/{id}`
- **Method**: `DELETE`
- **Description**: Delete a task by its ID.
- **Response**:
  ```json
  {
    "message": "Task deleted successfully."
  }
  ```

## Authentication

All endpoints require an API key passed in the headers as `Authorization: Bearer {API_KEY}`.

## Example Error Response

```json
{
  "error": "Error message describing the issue."
}
```

## Rate Limiting

The API supports rate limiting. Ensure your application adheres to the limit to avoid being blocked.

For more detailed information, please refer to the [official API documentation](https://documenter.getpostman.com/view/37344991/2sA3kdBHnF).
