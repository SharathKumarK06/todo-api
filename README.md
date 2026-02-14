# todo-golang
Todo application in golang

## Dependencies
- Web framework - `Gin`
- Database - `PostgreSQL`
- `Docker` for running database
- ORM - `GORM`
- API - `REST API`

### Installation of dependencies
- Database setup
```bash
docker run --name postgres \
  -e POSTGRES_USER=todo \
  -e POSTGRES_PASSWORD=pass \
  -e POSTGRES_DB=todo_db \
  -p 5432:5432 \
  -d postgres
```
- Install go dependencies
```bash
go mod tidy
```
- Check db/tables inside container
```bash
docker exec -it postgres psql -U todo -d todo_db
```

```sql
-- Show tables
\dt

-- Show entries of `todos` table
SELECT * FROM todos;
```

## Endpoints
```
GET     /todos
POST    /todos
PUT     /todos/:id
DELETE  /todos/:id
```



