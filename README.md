# todo-api
API for todo app in golang

## Dependencies
- Web framework - `Gin`
- Database - `PostgreSQL`
- `Docker` for running api and database
- ORM - `GORM`
- API - `REST API`

### Installation of dependencies
- Install `docker` and `docker-compose-v2`
- Setup full application
```bash
git clone https://github.com/SharathKumarK06/todo-api
cd todo-api
docker compose up --build -d
```

## Endpoints
```
GET     /todos          List todos
POST    /todos          Create new todo {"title": "Title of todo", "completed": false}
PUT     /todos/:id      Update todo {"title": "New title of todo", "completed": false}
DELETE  /todos/:id      Delete todo
```

## TODO
- docker compose setup
- Implemented database migrations and environment-based configuration.
- Tested API endpoints using curl.
- input validation and error handling for API endpoints
- logging and middleware for request handling.
- database schema for efficient task management
- Makefile or scripts to simplify development workflow.

